package handlers

import (
	"encoding/json"
	"net/http"
	"notes/model"
	"notes/store"
	"sync/atomic"
	"time"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var n model.Note
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	newId := atomic.AddInt64(&store.NextId, 1)
	n.ID = int(newId)
	now := time.Now()
	n.CreatedAt = now
	n.UpdatedAt = now
	store.Data = append(store.Data, n)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err1 := json.NewEncoder(w).Encode(n)
	if err1 != nil {
		http.Error(w, "invalid json", http.StatusInternalServerError)
		return
	}
}
