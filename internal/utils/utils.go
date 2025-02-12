package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"strings"
)

func GenerateShortUrl() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}
