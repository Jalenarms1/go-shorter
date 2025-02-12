package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Jalenarms1/go-shorter/internal/db"
	"github.com/Jalenarms1/go-shorter/internal/utils"
)

type URLRequest struct {
	URL string `json:"url"`
}

func HandleNewUrl(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	var req URLRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(req.URL)

	urlCode := utils.GenerateShortUrl()

	fmt.Println(urlCode)

	appUrl := db.AppUrl{
		SrcUrl:  req.URL,
		UrlCode: urlCode,
	}

	myDomain := os.Getenv("ROOT_DOMAIN")
	if myDomain == "" {
		http.Error(w, "Missing env - ROOT_DOMAIN", http.StatusBadRequest)
		return
	}

	if err := appUrl.Save(); err != nil {
		fmt.Println(err)
		http.Error(w, "Error saving new url", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"url": fmt.Sprintf("%s/go/%s", myDomain, appUrl.UrlCode)})

}
