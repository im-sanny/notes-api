package handlers

import (
	"encoding/json"
	"net/http"
	"notes/model"
	"notes/store"
	"strconv"
	"time"
)

func Put(w http.ResponseWriter, r *http.Request) {
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
			n.ID = id
			n.CreatedAt = store.Data[i].CreatedAt
			n.UpdatedAt = time.Now()
			store.Data[i] = n

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(store.Data[i])

			return
		}
	}
	http.Error(w, "note not found", http.StatusNotFound)
}
