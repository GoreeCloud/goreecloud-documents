package processing

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type PostgresQueue struct {
	DB  *sql.DB
	Now func() time.Time
}

func (q PostgresQueue) Claim(jobType JobType) (Job, error) {
	if q.DB == nil {
		return Job{}, ErrQueueUnavailable
	}
	row := q.DB.QueryRowContext(context.Background(), `
WITH candidate AS (
    SELECT id
    FROM processing_jobs
    WHERE state = 'queued' AND job_type = $1 AND available_at <= now()
    ORDER BY available_at, created_at, id
    FOR UPDATE SKIP LOCKED
    LIMIT 1
)
UPDATE processing_jobs AS jobs
SET state = 'running', attempts = attempts + 1, updated_at = now()
FROM candidate
WHERE jobs.id = candidate.id
RETURNING jobs.id::text, jobs.document_id::text, jobs.job_type, jobs.state, jobs.attempts, jobs.available_at, jobs.created_at, jobs.updated_at`, string(jobType))
	return scanJob(row)
}

func (q PostgresQueue) Complete(id string) (Job, error) { return q.transition(id, JobCompleted, 0) }
func (q PostgresQueue) Fail(id string) (Job, error)     { return q.transition(id, JobFailed, 0) }

func (q PostgresQueue) Retry(id string, delay time.Duration) (Job, error) {
	if q.DB == nil {
		return Job{}, ErrQueueUnavailable
	}
	if delay < 0 {
		return Job{}, fmt.Errorf("retry delay must not be negative")
	}
	row := q.DB.QueryRowContext(context.Background(), `
UPDATE processing_jobs
SET state = 'queued', available_at = now() + ($2 * interval '1 microsecond'), updated_at = now()
WHERE id = $1 AND state = 'running'
RETURNING id::text, document_id::text, job_type, state, attempts, available_at, created_at, updated_at`, id, delay.Microseconds())
	return scanJob(row)
}

func (q PostgresQueue) transition(id string, state JobState, _ time.Duration) (Job, error) {
	if q.DB == nil {
		return Job{}, ErrQueueUnavailable
	}
	row := q.DB.QueryRowContext(context.Background(), `
UPDATE processing_jobs
SET state = $2, updated_at = now()
WHERE id = $1 AND state = 'running'
RETURNING id::text, document_id::text, job_type, state, attempts, available_at, created_at, updated_at`, id, string(state))
	return scanJob(row)
}

type rowScanner interface{ Scan(...any) error }

func scanJob(row rowScanner) (Job, error) {
	var job Job
	var typ, state string
	if err := row.Scan(&job.ID, &job.DocumentID, &typ, &state, &job.Attempts, &job.Available, &job.CreatedAt, &job.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Job{}, ErrNoAvailableJob
		}
		return Job{}, fmt.Errorf("processing queue: %w", err)
	}
	job.Type = JobType(typ)
	job.State = JobState(state)
	if err := job.Validate(); err != nil {
		return Job{}, fmt.Errorf("processing queue returned invalid job: %w", err)
	}
	return job, nil
}
