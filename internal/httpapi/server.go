// Package httpapi defines the initial GoreeCloud Documents HTTP boundary.
package httpapi

import (
	"encoding/json"
	"net/http"
	"time"
)

const APIPath = "/api/v1"

// Server exposes only development-safe foundation endpoints until production
// authentication and persistent repositories are implemented.
type Server struct {
	startedAt time.Time
	mux       *http.ServeMux
}

// New constructs the initial fail-closed API surface.
func New() *Server {
	s := &Server{startedAt: time.Now().UTC(), mux: http.NewServeMux()}
	s.mux.HandleFunc("GET /healthz", s.health)
	s.mux.HandleFunc("GET "+APIPath+"/status", s.status)
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Referrer-Policy", "no-referrer")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	s.mux.ServeHTTP(w, r)
}

func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status":  "ok",
		"service": "goreecloud-documents",
	})
}

func (s *Server) status(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"service":       "goreecloud-documents",
		"lifecycle":     "development",
		"api_version":   "v1",
		"native_product": true,
		"started_at":    s.startedAt,
		"capabilities": []string{
			"domain-foundation",
			"processing-contracts",
			"ocr-adapter-contract",
			"search-adapter-contract",
		},
	})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
