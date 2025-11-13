package router

import (
	"hafiztri123/be-sharing-vision/internal/articles"
	"net/http"
)

func NewRouter(h *articles.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/article/{limit}/{offset}", h.GetArticlesHandler)
	mux.HandleFunc("GET /api/v1/article/{id}", h.GetArticleHandler)
	mux.HandleFunc("PUT /api/v1/article/{id}", h.UpdateArticleHandler)
	mux.HandleFunc("DELETE /api/v1/article/{id}", h.DeleteArticleHandler)
	mux.HandleFunc("POST /api/v1/article", h.CreateArticleHandler)

	return mux
}
