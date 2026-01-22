package handlers

import (
	"blog/internal/storage"
	"encoding/json"
	"net/http"
)

func GetAllBlogs(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Blog)

}
