package handlers

import (
	"encoding/json"
	"net/http"
	"notes/model"
	"notes/store"
	"strconv"
	"time"
)

func Patch(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var n model.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, "invalid json format", http.StatusBadRequest)
		return
	}

	for i := range store.Data {
		if store.Data[i].ID == id {
			if n.Title != "" {
				store.Data[i].Title = n.Title
			}
			if n.Content != "" {
				store.Data[i].Content = n.Content
			}
			n.UpdatedAt = time.Now()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(store.Data[i])
			return
		}
	}
}
