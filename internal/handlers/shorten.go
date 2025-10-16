package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var short_url string
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
			return
		}

		url_original := r.FormValue("url")

		if short_url, err = h.repo.GetByOriginalURL(url_original); err != nil {
			if err == sql.ErrNoRows {
				short_url = h.createShortURL(url_original)
				fmt.Fprint(w, "Your url:", url_original)
				fmt.Fprint(w, "\nShorten url: ", short_url)
			} else {
				log.Fatalf("Error ShortenURL: %v", err)
			}
		} else {
			fmt.Fprint(w, "Your url: ", url_original)
			fmt.Fprint(w, "\nShorten url: ", short_url)
		}
	} else {
		http.Error(w, "link not found", http.StatusServiceUnavailable)
		return
	}
}

func (h *Handler) createShortURL(url string) string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	short_url := "http://localhost:8080/" + hex.EncodeToString(bytes)[:5]
	err := h.repo.Insert(short_url, url)
	fmt.Println(short_url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Запись успешно записана!")

	return short_url
}
