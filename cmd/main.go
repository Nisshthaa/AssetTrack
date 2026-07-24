package main

import (
	"AssetTrack/database"
	"AssetTrack/server"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const shutDownTimeOut = 10 * time.Second

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := server.SetUpRoutes()

	if err := database.Connect(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")); err != nil {
		log.Panicf("failed to initialize and migrate database with error: %+v", err)
	}

	log.Print("migration successful")

	go func() {
		if serverErr := srv.Run(":8080"); serverErr != nil && !errors.Is(serverErr, http.ErrServerClosed) {
			log.Panicf("Failed to run server with error: %+v", serverErr)
		}
	}()

	log.Print("server started at :8080")

	<-done

	log.Println("shutting down server")

	if serverCloseErr := srv.Shutdown(shutDownTimeOut); serverCloseErr != nil {
		log.Panicf("failed to gracefully shutdown server: %v", serverCloseErr)
	}

	if dbCloseErr := database.ShutdownDatabase(); dbCloseErr != nil {
		log.Println("failed to close database connection:", dbCloseErr)
	}

}
