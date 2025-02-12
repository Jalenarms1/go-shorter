package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jalenarms1/go-shorter/internal/db"
)

func main() {
	// if godotenv.Load() != nil {
	// 	log.Fatal("ENV not loaded")
	// }

	if !db.SetDB() {
		log.Fatal("DB not set")
	}

	fmt.Println("DB Connected")

	mux := NewServer()

	log.Fatal(http.ListenAndServe(":8080", mux))
}
