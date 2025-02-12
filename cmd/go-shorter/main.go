package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jalenarms1/go-shorter/internal/db"
	"github.com/Jalenarms1/go-shorter/internal/utils"
)

func main() {
	// if godotenv.Load() != nil {
	// 	log.Fatal("ENV not loaded")
	// }
	urlHash := utils.GenerateShortUrl()

	fmt.Println(urlHash)

	if !db.SetDB() {
		log.Fatal("DB not set")
	}

	fmt.Println("DB Connected")

	mux := NewServer()

	log.Fatal(http.ListenAndServe(":8080", mux))
}
