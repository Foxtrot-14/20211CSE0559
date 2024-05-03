package router

import (
	"backend/internal/controller"
	"net/http"
)
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/{categoryname}/products/{n}",controller.CategoryNameHandler)
	return mux
}