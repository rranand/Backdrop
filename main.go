package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rranand/backdrop/prisma/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		fmt.Println("Error Occurred : " + err.Error())
	}

	fmt.Println("Connection Created!")
	client.Disconnect()

}

// func sample_api() {
// 	http.HandleFunc("/", sampleFunc)
// 	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
// 	http.ListenAndServe(addr, nil)
// }

// func sampleFunc(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "Hello, From Go\n")
// }
