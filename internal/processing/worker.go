package processing

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Queue is the state-transition authority used by workers. Durable queues must
// preserve the same atomic claim and running-only transition semantics.
type Queue interface {
	Claim(JobType) (Job, error)
	Complete(string) (Job, error)
	Retry(string, time.Duration) (Job, error)
	Fail(string) (Job, error)
}

// Handler performs one bounded processing operation. It does not own job state.
type Handler interface {
	Process(context.Context, Job) error
}

type Worker struct {
	Queue      Queue
	Handler    Handler
	JobType    JobType
	RetryDelay time.Duration
	MaxAttempts int
}

// RunOnce claims at most one job and records the outcome. A handler failure is
// retried only while the configured attempt budget remains; otherwise it is
// terminally failed. Context cancellation never converts a job to success.
func (w Worker) RunOnce(ctx context.Context) (Job, error) {
	if w.Queue == nil || w.Handler == nil || w.MaxAttempts <= 0 || w.RetryDelay < 0 {
		return Job{}, fmt.Errorf("processing worker unavailable")
	}
	job, err := w.Queue.Claim(w.JobType)
	if err != nil {
		return Job{}, err
	}
	if err := w.Handler.Process(ctx, job); err != nil {
		if job.Attempts < w.MaxAttempts && !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
			updated, transitionErr := w.Queue.Retry(job.ID, w.RetryDelay)
			if transitionErr != nil {
				return Job{}, fmt.Errorf("record processing retry: %w", transitionErr)
			}
			return updated, err
		}
		updated, transitionErr := w.Queue.Fail(job.ID)
		if transitionErr != nil {
			return Job{}, fmt.Errorf("record processing failure: %w", transitionErr)
		}
		return updated, err
	}
	updated, err := w.Queue.Complete(job.ID)
	if err != nil {
		return Job{}, fmt.Errorf("record processing completion: %w", err)
	}
	return updated, nil
}
