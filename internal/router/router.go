package router

import (
	"blog/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetUpRouter() http.Handler{
	router:= chi.NewRouter()
	router.Route("/blogs", func(r chi.Router){
		r.Get("/", handlers.GetAllBlogs)
	}),
}
