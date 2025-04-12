package database

import (
	"context"
	"log"

	"github.com/rranand/backdrop/prisma/db"
)

var Client *db.PrismaClient

func Connect() error {
	Client = db.NewClient()

	if err := Client.Prisma.Connect(); err != nil {
		log.Fatal("Error Occurred : " + err.Error())
		return err
	}

	log.Println("DB Connection Created!")
	return nil
}

func Disconnect(ctx context.Context) {
	if Client == nil {
		log.Println("⚠️ DB client is nil, nothing to disconnect")
		return
	}
	done := make(chan error, 1)
	go func() {
		log.Println("🔌 Attempting to disconnect DB...")
		done <- Client.Prisma.Disconnect()
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
