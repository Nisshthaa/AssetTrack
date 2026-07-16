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

func SetUpRoutes() *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/register", handlers.RegisterUser)
	mux.HandleFunc("POST /v1/login", handlers.LoginUser)

	mux.Handle("GET /v1/profile", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.GetUser), "admin", "project_manager", "employee")))
	mux.Handle("GET /v1/my-assets", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.GetUserAssets), "employee")))
	mux.Handle("GET /v1/my-asset", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.GetUserAssetByID), "employee")))

	mux.Handle("POST /v1/assets", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.CreateAsset), "admin")))
	mux.Handle("GET /v1/assets", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.GetAssets), "admin", "project_manager")))
	mux.Handle("GET /v1/assets/{assetID}", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.GetAssetByID), "admin", "project_manager")))
	mux.Handle("PUT /v1/assets/{assetID}", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.UpdateAsset), "admin")))
	mux.Handle("POST /v1/assets/{assetID}/repair", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.AssetSentToRepair), "admin")))
	mux.Handle("POST /v1/assets/{assetID}/repaired", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.AssetRepairCompleted), "admin")))
	mux.Handle("DELETE /v1/assets/{assetID}", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.DeleteAsset), "admin")))

	mux.Handle("POST /v1/assets/assign", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.AssignAsset), "admin", "project_manager")))
	mux.Handle("PUT /v1/assets/return", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.ReturnAsset), "admin", "project_manager")))
	mux.Handle("GET /v1/assets/{assetID}/history", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.GetAssetHistory), "admin")))

	mux.Handle("GET /v1/dashboard", middlewares.Authenticate(middlewares.RequireRoles(http.HandlerFunc(handlers.AdminDashboard), "admin")))

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
