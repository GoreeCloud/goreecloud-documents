package processing

import (
	"errors"
	"testing"
	"time"
)

func TestMemoryQueueClaimIsSingleAndCompletes(t *testing.T) {
	now := time.Date(2026, 8, 29, 20, 0, 0, 0, time.UTC)
	queue := NewMemoryQueue()
	queue.now = func() time.Time { return now }
	job := Job{ID: "job-1", DocumentID: "doc-1", Type: JobOCR, State: JobQueued, Available: now, CreatedAt: now, UpdatedAt: now}
	if err := queue.Enqueue(job); err != nil {
		t.Fatal(err)
	}
	claimed, err := queue.Claim(JobOCR)
	if err != nil {
		t.Fatal(err)
	}
	if claimed.State != JobRunning || claimed.Attempts != 1 {
		t.Fatalf("claimed=%#v", claimed)
	}
	if _, err := queue.Claim(JobOCR); !errors.Is(err, ErrJobNotFound) {
		t.Fatalf("second claim err=%v", err)
	}
	completed, err := queue.Complete(claimed.ID)
	if err != nil {
		t.Fatal(err)
	}
	if completed.State != JobCompleted {
		t.Fatalf("state=%s", completed.State)
	}
}

func TestMemoryQueueRetryHonorsAvailability(t *testing.T) {
	now := time.Date(2026, 8, 29, 20, 0, 0, 0, time.UTC)
	queue := NewMemoryQueue()
	queue.now = func() time.Time { return now }
	job := Job{ID: "job-1", DocumentID: "doc-1", Type: JobIndex, State: JobQueued, Available: now, CreatedAt: now, UpdatedAt: now}
	if err := queue.Enqueue(job); err != nil {
		t.Fatal(err)
	}
	claimed, _ := queue.Claim(JobIndex)
	if _, err := queue.Retry(claimed.ID, time.Minute); err != nil {
		t.Fatal(err)
	}
	if _, err := queue.Claim(JobIndex); !errors.Is(err, ErrJobNotFound) {
		t.Fatalf("early claim err=%v", err)
	}
	queue.now = func() time.Time { return now.Add(time.Minute) }
	reclaimed, err := queue.Claim(JobIndex)
	if err != nil {
		t.Fatal(err)
	}
	if reclaimed.Attempts != 2 {
		t.Fatalf("attempts=%d", reclaimed.Attempts)
	}
}
