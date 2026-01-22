package handlers

import (
	"encoding/json"
	"net/http"
	"notes/store"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(store.Data)
	if err != nil {
		http.Error(w, "invalid json", http.StatusInternalServerError)
		return
	}
}
