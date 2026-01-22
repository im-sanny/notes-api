package handlers

import (
	"encoding/json"
	"net/http"
	"notes/store"
	"strconv"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for _, v := range store.Data {
		if v.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(v)
			if err != nil {
				http.Error(w, "invalid json", http.StatusInternalServerError)
				return
			}
			return
		}
	}
	http.Error(w, "note not found", http.StatusNotFound)
}
