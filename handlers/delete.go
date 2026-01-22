package handlers

import (
	"net/http"
	"notes/store"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for i := range store.Data {
		if store.Data[i].ID == id {
			store.Data = append(store.Data[0:i], store.Data[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "note not found", http.StatusNotFound)
}
