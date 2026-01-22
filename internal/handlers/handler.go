package handlers

import (
	"blog/internal/storage"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAllBlogs(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storage.Blog)
}

func GetBlogByID(w http.ResponseWriter, r *http.Request){
	id, err:=strconv.Atoi(chi.URLParam(r, "id"))
	if(err!=nil){
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _,blog := range(storage.Blog){
		if(blog.ID==id){
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(blog)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func CreateBlog(){

}

func UpdateBlog(){

}

func DeleteBlog(){

}
