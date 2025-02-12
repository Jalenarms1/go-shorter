package db

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func SetDB() bool {
	url := os.Getenv("TURSO_DATABASE_URL")
	token := os.Getenv("TURSO_AUTH_TOKEN")

	if url == "" || token == "" {
		log.Fatal("Missing db url or token")

	}

	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	http.DefaultTransport = customTransport

	var err error
	DB, err = sql.Open("libsql", fmt.Sprintf("%s?authToken=%s", url, token))

	return err == nil

}
