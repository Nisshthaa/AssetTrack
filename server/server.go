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

func protectedWithRoles(h http.HandlerFunc, roles ...string) http.Handler {
	return middlewares.Authenticate(
		middlewares.RequireRoles(roles...)(h),
	)
}

func SetUpRoutes() *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/register", handlers.RegisterUser)
	mux.HandleFunc("POST /v1/login", handlers.LoginUser)

	mux.Handle("GET /v1/profile", protectedWithRoles(handlers.GetUser, "admin", "project_manager", "employee"))
	mux.Handle("POST /v1/logout", protectedWithRoles(handlers.LogoutUser, "admin", "project_manager", "employee"))
	mux.Handle("DELETE /v1/delete", protectedWithRoles(handlers.DeleteUser, "admin", "project_manager", "employee"))

	mux.Handle("POST /v1/assets", protectedWithRoles(handlers.CreateAsset, "admin"))
	mux.Handle("GET /v1/assets", protectedWithRoles(handlers.GetAssets, "admin", "project_manager"))
	mux.Handle("GET /v1/assets/{assetID}", protectedWithRoles(handlers.GetAssetByID, "admin", "project_manager"))
	mux.Handle("PUT /v1/assets/{assetID}", protectedWithRoles(handlers.UpdateAsset, "admin"))
	mux.Handle("DELETE /v1/assets/{assetID}", protectedWithRoles(handlers.DeleteAsset, "admin"))

	mux.Handle("POST /v1/assets/assign", protectedWithRoles(handlers.AssignAsset, "admin", "project_manager"))
	mux.Handle("PUT /v1/assets/return", protectedWithRoles(handlers.ReturnAsset, "admin", "project_manager"))

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
