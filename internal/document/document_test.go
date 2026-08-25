package document

import (
	"strings"
	"testing"
	"time"
)

func validRecord() Record {
	now := time.Now().UTC()
	return Record{
		ID:              "11111111-1111-4111-8111-111111111111",
		OwnerID:         "22222222-2222-4222-8222-222222222222",
		Title:           "Example record",
		OriginalName:    "example.pdf",
		MediaType:       "application/pdf",
		StorageKey:      "documents/11111111-1111-4111-8111-111111111111/original",
		ChecksumSHA256:  strings.Repeat("a", 64),
		SizeBytes:       1024,
		ProcessingState: StatePending,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func TestRecordValidate(t *testing.T) {
	record := validRecord()
	if err := record.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestRecordValidateRejectsInvalidIdentityAndChecksum(t *testing.T) {
	record := validRecord()
	record.OwnerID = "../owner"
	if err := record.Validate(); err == nil {
		t.Fatal("Validate() unexpectedly accepted invalid owner ID")
	}

	record = validRecord()
	record.ChecksumSHA256 = "not-a-digest"
	if err := record.Validate(); err == nil {
		t.Fatal("Validate() unexpectedly accepted invalid checksum")
	}
}
