// Package ocr defines the supporting-component boundary for OCR engines.
package ocr

import "context"

// Request contains only the information an OCR provider needs to process one
// preserved source document. Authentication and authorization are resolved
// before this boundary is invoked.
type Request struct {
	DocumentID string
	MediaType  string
	SourcePath string
	Languages  []string
}

// Result contains provider output without making any OCR engine the authority
// for GoreeCloud Documents metadata or document lifecycle.
type Result struct {
	Text       string
	PageCount  int
	Language   string
	Confidence float64
}

// Provider is implemented by approved OCR supporting components.
type Provider interface {
	Name() string
	Recognize(context.Context, Request) (Result, error)
}
