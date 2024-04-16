package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func BookRoutes() chi.Router{
	r:= chi.NewRouter()
	BookHandler:= BookHandler{}
	r.Get("/", BookHandler.ListBooks)
	r.Post("/CreateBook", BookHandler.DeleteBook)
	r.Get("/{id}", BookHandler.GetBooks)
	r.Put("/{id}", BookHandler.UpdateBook)
	r.Delete("/{id}", BookHandler.DeleteBook)
	return r
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Mount("/books", BookRoutes())

	http.ListenAndServe(":3001", r)
}
