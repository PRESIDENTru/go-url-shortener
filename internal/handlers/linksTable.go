package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) Links(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		links, err := h.repo.GetAll()
		if err != nil {
			http.Error(w, "DB Ошибка", http.StatusInternalServerError)
			return
		}

		for _, l := range links {
			fmt.Fprintln(w, "id: ", l.ID)
			fmt.Fprintln(w, "short_url: ", l.ShortURL)
			fmt.Fprintln(w, "original_url: ", l.OriginalURL)
			fmt.Fprintln(w, "time_creaded: ", l.TimeCreated)
			fmt.Fprintln(w, "-------------------------------------------------------------")
		}
	} else {
		http.Error(w, "link not found", http.StatusServiceUnavailable)
		return
	}
}
