package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Jalenarms1/go-shorter/internal/db"
	"github.com/Jalenarms1/go-shorter/internal/utils"
)

type URLRequest struct {
	URL string `json:"url"`
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	urlCode := r.URL.Path[len("/go/"):]

	appUrl, err := db.GetAppUrl(urlCode)
	if err != nil {
		http.Error(w, "Url not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, appUrl.SrcUrl, http.StatusFound)
}

func HandleNewUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	if !isValidUrl(req.URL) {
		http.Error(w, "Invalid URL provided", http.StatusBadRequest)
		return
	}

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

func isValidUrl(srcUrl string) bool {
	parsedUrl, err := url.ParseRequestURI(srcUrl)
	if err != nil {
		return false
	}

	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return false
	}

	client := http.Client{}

	resp, err := client.Get(srcUrl)
	if err != nil {
		fmt.Print(err)
		return false
	}

	if resp.StatusCode != 200 {
		return false
	}

	return true
}
