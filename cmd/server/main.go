package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/rranand/backdrop/internal/router"
	"github.com/rranand/backdrop/pkg/constants"
	"github.com/rranand/backdrop/pkg/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Server StartUp Failed, Error while loading .env file")
		return
	}

	err = database.Connect(os.Getenv("DB_URL"))

	if err != nil {
		log.Fatal("Server StartUp Failed, Error while Connection To DB")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), constants.DatabaseConnectionTimeoutDuration*time.Second)
	defer cancel()
	defer database.Disconnect(ctx)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	srv := &http.Server{
		Addr:    addr,
		Handler: router.Router(),
	}

	go func() {
		log.Printf("Server started at %s. Press Cmd+Z to disconnect DB and exit...", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGSEGV)
	<-stop

	ctx_force_exit, cancel_force_exit := context.WithTimeout(context.Background(), constants.ServerStopForcefulTimeoutDuration)
	defer cancel_force_exit()

	if err := srv.Shutdown(ctx_force_exit); err != nil {
		log.Fatalf("❌ Server shutdown failed: %v", err)
	}

	database.Disconnect(ctx_force_exit)
	log.Println("👋 Server and DB shut down cleanly")

}

//nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/server/main.go
//go run cmd/server/main.go
