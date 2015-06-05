package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const API_URL = "https://api.github.com/repos/%s/releases/latest"

func main() {
	http.HandleFunc("/", handler)

	port := "8888"

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	log.Println("Waiting for requests on", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	str := strings.Split(r.URL.Path[1:], "/")
	if len(str) < 3 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	user := str[0]
	repo := str[1]
	file := str[2]
	url := fmt.Sprintf(API_URL, user+"/"+repo)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var obj map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var assets []interface{}
	var ok bool
	if assets, ok = obj["assets"].([]interface{}); !ok {
		http.Error(w, "Invalid response from GitHub", http.StatusInternalServerError)
		return
	}

	for _, asset := range assets {
		if m, ok := asset.(map[string]interface{}); ok {
			if name, ok := m["name"].(string); ok {
				if name == file {
					if dUrl, ok := m["browser_download_url"].(string); ok {
						log.Println(r.URL.Path, r.Host, dUrl)
						http.Redirect(w, r, dUrl, http.StatusFound)
						return
					}
				}
			}
		}
	}

	http.Error(w, "Requested file not found", http.StatusNotFound)
}
