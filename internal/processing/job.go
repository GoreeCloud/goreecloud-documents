// Package processing defines durable ingestion and processing work contracts.
package processing

import (
	"fmt"
	"time"
)

// JobType identifies one stage of the document-processing pipeline.
type JobType string

const (
	JobIngest  JobType = "ingest"
	JobOCR     JobType = "ocr"
	JobIndex   JobType = "index"
	JobExtract JobType = "extract"
)

// JobState records durable worker progress.
type JobState string

const (
	JobQueued    JobState = "queued"
	JobRunning   JobState = "running"
	JobCompleted JobState = "completed"
	JobFailed    JobState = "failed"
)

// Job represents bounded, retryable work for a single document.
type Job struct {
	ID         string    `json:"id"`
	DocumentID string    `json:"document_id"`
	Type       JobType   `json:"type"`
	State      JobState  `json:"state"`
	Attempts   int       `json:"attempts"`
	Available  time.Time `json:"available_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Validate rejects invalid job state before persistence or dispatch.
func (j Job) Validate() error {
	if j.ID == "" || j.DocumentID == "" {
		return fmt.Errorf("job and document IDs must not be empty")
	}
	switch j.Type {
	case JobIngest, JobOCR, JobIndex, JobExtract:
	default:
		return fmt.Errorf("unsupported job type %q", j.Type)
	}
	switch j.State {
	case JobQueued, JobRunning, JobCompleted, JobFailed:
	default:
		return fmt.Errorf("unsupported job state %q", j.State)
	}
	if j.Attempts < 0 {
		return fmt.Errorf("job attempts must not be negative")
	}
	if j.CreatedAt.IsZero() || j.UpdatedAt.IsZero() || j.UpdatedAt.Before(j.CreatedAt) {
		return fmt.Errorf("job timestamps are invalid")
	}
	return nil
}
