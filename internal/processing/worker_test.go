package processing

import (
	"context"
	"errors"
	"testing"
	"time"
)

type handlerFunc func(context.Context, Job) error

func (f handlerFunc) Process(ctx context.Context, job Job) error { return f(ctx, job) }

func queuedJob() Job {
	now := time.Now().UTC()
	return Job{
		ID:         "11111111-1111-1111-1111-111111111111",
		DocumentID: "22222222-2222-2222-2222-222222222222",
		Type:       JobOCR,
		State:      JobQueued,
		Available:  now.Add(-time.Minute),
		CreatedAt:  now.Add(-time.Minute),
		UpdatedAt:  now.Add(-time.Minute),
	}
}

func TestWorkerCompletesSuccessfulJob(t *testing.T) {
	queue := NewMemoryQueue()
	if err := queue.Enqueue(queuedJob()); err != nil {
		t.Fatal(err)
	}
	worker := Worker{
		Queue:       queue,
		Handler:     handlerFunc(func(context.Context, Job) error { return nil }),
		JobType:     JobOCR,
		RetryDelay:  time.Second,
		MaxAttempts: 3,
	}
	job, err := worker.RunOnce(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if job.State != JobCompleted || job.Attempts != 1 {
		t.Fatalf("job=%+v", job)
	}
}

func TestWorkerRetriesWithinAttemptBudget(t *testing.T) {
	queue := NewMemoryQueue()
	if err := queue.Enqueue(queuedJob()); err != nil {
		t.Fatal(err)
	}
	boom := errors.New("ocr unavailable")
	worker := Worker{
		Queue:       queue,
		Handler:     handlerFunc(func(context.Context, Job) error { return boom }),
		JobType:     JobOCR,
		RetryDelay:  time.Second,
		MaxAttempts: 3,
	}
	job, err := worker.RunOnce(context.Background())
	if !errors.Is(err, boom) {
		t.Fatalf("err=%v", err)
	}
	if job.State != JobQueued || job.Attempts != 1 {
		t.Fatalf("job=%+v", job)
	}
}

func TestWorkerFailsAtAttemptBudget(t *testing.T) {
	queue := NewMemoryQueue()
	job := queuedJob()
	job.Attempts = 1
	if err := queue.Enqueue(job); err != nil {
		t.Fatal(err)
	}
	boom := errors.New("ocr failed")
	worker := Worker{
		Queue:       queue,
		Handler:     handlerFunc(func(context.Context, Job) error { return boom }),
		JobType:     JobOCR,
		RetryDelay:  0,
		MaxAttempts: 2,
	}
	updated, err := worker.RunOnce(context.Background())
	if !errors.Is(err, boom) {
		t.Fatalf("err=%v", err)
	}
	if updated.State != JobFailed || updated.Attempts != 2 {
		t.Fatalf("job=%+v", updated)
	}
}
