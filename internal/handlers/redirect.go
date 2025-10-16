package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		short_url := "http://localhost:8080" + r.URL.Path
		fmt.Println(short_url)
		var original_url string

		original_url, err := h.repo.GetByShortURL(short_url)
		if err != nil {
			http.Error(w, "Link not found", http.StatusNotFound)
			return
		}
		fmt.Println("Ссылка найдена")
		http.Redirect(w, r, original_url, http.StatusFound)
	} else {
		http.Error(w, "link not found", http.StatusServiceUnavailable)
		return
	}

}
