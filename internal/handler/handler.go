package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WebRouter struct {
	*mux.Router
}

func NewWebRouter() *WebRouter {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./web"))
	r.PathPrefix("/").Handler(fs)

	return &WebRouter{
		Router: r,
	}
}
