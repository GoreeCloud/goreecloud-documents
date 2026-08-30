package processing

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrJobNotFound = errors.New("processing job not found")

// MemoryQueue is a deterministic development queue for bounded worker
// orchestration. Durable PostgreSQL dispatch can implement the same semantics.
type MemoryQueue struct {
	mu   sync.Mutex
	jobs map[string]Job
	now  func() time.Time
}

func NewMemoryQueue() *MemoryQueue {
	return &MemoryQueue{jobs: make(map[string]Job), now: time.Now}
}

func (q *MemoryQueue) Enqueue(job Job) error {
	if q == nil {
		return fmt.Errorf("processing queue unavailable")
	}
	if err := job.Validate(); err != nil {
		return err
	}
	if job.State != JobQueued {
		return fmt.Errorf("new processing job must be queued")
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	if _, exists := q.jobs[job.ID]; exists {
		return fmt.Errorf("processing job already exists")
	}
	q.jobs[job.ID] = job
	return nil
}

// Claim returns the oldest available queued job of the requested type and marks
// it running atomically so two workers cannot claim the same in-memory job.
func (q *MemoryQueue) Claim(jobType JobType) (Job, error) {
	if q == nil {
		return Job{}, fmt.Errorf("processing queue unavailable")
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	now := q.now().UTC()
	var selected *Job
	for _, candidate := range q.jobs {
		if candidate.Type != jobType || candidate.State != JobQueued || candidate.Available.After(now) {
			continue
		}
		candidateCopy := candidate
		if selected == nil || candidateCopy.CreatedAt.Before(selected.CreatedAt) {
			selected = &candidateCopy
		}
	}
	if selected == nil {
		return Job{}, ErrJobNotFound
	}
	selected.State = JobRunning
	selected.Attempts++
	selected.UpdatedAt = now
	q.jobs[selected.ID] = *selected
	return *selected, nil
}

func (q *MemoryQueue) Complete(jobID string) (Job, error) {
	return q.transition(jobID, JobCompleted, time.Time{})
}

// Retry returns a running job to queued state after a bounded delay.
func (q *MemoryQueue) Retry(jobID string, delay time.Duration) (Job, error) {
	if delay < 0 {
		return Job{}, fmt.Errorf("retry delay must not be negative")
	}
	return q.transition(jobID, JobQueued, q.now().UTC().Add(delay))
}

func (q *MemoryQueue) Fail(jobID string) (Job, error) {
	return q.transition(jobID, JobFailed, time.Time{})
}

func (q *MemoryQueue) transition(jobID string, state JobState, available time.Time) (Job, error) {
	if q == nil {
		return Job{}, fmt.Errorf("processing queue unavailable")
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	job, ok := q.jobs[jobID]
	if !ok {
		return Job{}, ErrJobNotFound
	}
	if job.State != JobRunning {
		return Job{}, fmt.Errorf("processing job is not running")
	}
	job.State = state
	if state == JobQueued {
		job.Available = available
	}
	job.UpdatedAt = q.now().UTC()
	q.jobs[jobID] = job
	return job, nil
}
