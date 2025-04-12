package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rranand/backdrop/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// client := db.NewClient()

	// if err := client.Prisma.Connect(); err != nil {
	// 	log.Fatal("Error Occurred : " + err.Error())
	// }

	// log.Println("Connection Created!")
	// client.Disconnect()

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Printf("Server started at %s", addr)
	http.ListenAndServe(addr, router.Router())
}

//nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/server/main.go

// func sample_api() {
// 	http.HandleFunc("/", sampleFunc)
// 	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
// 	http.ListenAndServe(addr, nil)
// }

// func sampleFunc(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "Hello, From Go\n")
// }
