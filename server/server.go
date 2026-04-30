// Package server provides the HTTP server setup and routing for gascity.
package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gastownhall/gascity/config"
)

// Server wraps the standard HTTP server with additional configuration
// and lifecycle management.
type Server struct {
	httpServer *http.Server
	cfg        *config.Config
	logger     *log.Logger
}

// New creates a new Server instance with the provided configuration.
// It sets up the router and applies all middleware.
func New(cfg *config.Config) *Server {
	logger := log.New(os.Stdout, "[gascity] ", log.LstdFlags|log.Lshortfile)

	mux := http.NewServeMux()
	s := &Server{
		cfg:    cfg,
		logger: logger,
	}

	s.registerRoutes(mux)

	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      s.applyMiddleware(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return s
}

// Start begins listening for incoming HTTP requests.
// It blocks until a shutdown signal is received, then performs a graceful shutdown.
func (s *Server) Start() error {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)\n
	errCh := make(chan error, 1)
	go func() {
		s.logger.Printf("server listening on %s", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("server error: %w", err)
		}
	}()

	select {
	case err := <-errCh:
		return err
	case sig := <-shutdownCh:
		s.logger.Printf("received signal %s, shutting down gracefully", sig)
		return s.Shutdown()
	}
}

// Shutdown gracefully stops the HTTP server, allowing in-flight requests
// up to 10 seconds to complete before forcefully closing.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("graceful shutdown failed: %w", err)
	}

	s.logger.Println("server stopped")
	return nil
}

// registerRoutes sets up all HTTP routes on the provided mux.
func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/", s.handleNotFound)
}

// applyMiddleware wraps the handler with common middleware such as
// request logging and recovery from panics.
func (s *Server) applyMiddleware(next http.Handler) http.Handler {
	return s.recoveryMiddleware(s.loggingMiddleware(next))
}

// loggingMiddleware logs each incoming request with method, path, and duration.
func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		s.logger.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// recoveryMiddleware catches any panics during request handling and
// returns a 500 Internal Server Error response.
func (s *Server) recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				s.logger.Printf("panic recovered: %v", rec)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// handleHealth responds to health check requests with a 200 OK status.
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

// handleNotFound returns a 404 for any unmatched routes.
func (s *Server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
