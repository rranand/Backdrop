package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func New(addr string) error {
	DB, err := sql.Open("postgres", addr)
	if err != nil {
		return err
	}

	maxOpenConns := 10
	maxIdleConns := 10
	maxIdleTime := "15m"

	DB.SetMaxOpenConns(maxOpenConns)
	DB.SetMaxIdleConns(maxIdleConns)

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return err
	}
	DB.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = DB.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func Connect(addr string) error {
	err := error(nil)
	DB, err = sql.Open("postgres", addr)
	if err != nil {
		log.Fatal("Error Occurred : " + err.Error())
		return err
	}

	log.Println("DB Connection Created!")
	return nil
}

func Disconnect(ctx context.Context) {
	if DB == nil {
		log.Println("⚠️ DB client is nil, nothing to disconnect")
		return
	}
	done := make(chan error, 1)
	go func() {
		log.Println("🔌 Attempting to disconnect DB...")
		done <- DB.Close()
	}()

	select {
	case err := <-done:
		if err != nil {
			log.Printf("Error disconnecting DB: %v", err)
		} else {
			log.Println("DB disconnected")
		}
	case <-ctx.Done():
		log.Println("Timed out during DB disconnect")
	}
}
