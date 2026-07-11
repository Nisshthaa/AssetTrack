package server

import (
	"AssetTrack/handlers"
	"AssetTrack/middlewares"
	"context"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	router http.Handler
}

const (
	readTimeout       = 5 * time.Minute
	readHeaderTimeout = 30 * time.Second
	writeTimeout      = 5 * time.Minute
)

func protected(h http.HandlerFunc) http.Handler {
	return middlewares.Authenticate(h)
}

func SetUpRoutes() *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/register", handlers.RegisterUser)
	mux.HandleFunc("POST /v1/login", handlers.LoginUser)

	mux.Handle("GET /v1/profile", protected(handlers.GetUser))

	return &Server{
		router: mux,
	}
}

func (s *Server) Run(port string) error {
	s.server = &http.Server{
		Addr:              port,
		Handler:           s.router,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
	}
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}
