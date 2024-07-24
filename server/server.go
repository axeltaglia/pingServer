package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// Server represents the HTTP server
type Server struct {
	serverMux *http.ServeMux
	port      string
}

// NewServer creates a new Server with the given port
func NewServer(port string) Server {
	serverMux := http.NewServeMux()
	return Server{
		serverMux: serverMux,
		port:      port,
	}
}

// Run starts the HTTP server and listens on the specified port
func (o *Server) Run() error {
	// Listen and serve on the specified port
	if err := http.ListenAndServe(o.port, o.serverMux); err != nil {
		return err
	}

	return nil
}

// HandleEndpoints sets up the endpoints for the server
func (o *Server) HandleEndpoints() {
	// Set up the /ping endpoint
	o.serverMux.HandleFunc("/ping", makeHttpFunc(ping))
}

// WebError represents an error that can be sent as a JSON response
type WebError struct {
	ErrMsg string `json:"errMsg"`
}

// makeHttpFunc wraps an HTTP handler function with error handling
func makeHttpFunc(f func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			slog.Error("Handler error: ", err)
			payload := WebError{ErrMsg: err.Error()}
			if err := writeJSON(w, http.StatusBadRequest, payload); err != nil {
				slog.Error("Couldn't write error response: ", err)
				return
			}
		}
	}
}

// writeJSON writes a JSON response to the client
func writeJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// PingResponse represents the response for the /ping endpoint
type PingResponse struct {
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

// ping handles the /ping endpoint and returns a JSON response
func ping(w http.ResponseWriter, _ *http.Request) error {
	response := PingResponse{
		Id:  0,
		Msg: "Pong",
	}
	return writeJSON(w, http.StatusOK, response)
}
