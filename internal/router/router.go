package router

import (
	"hafiztri123/be-sharing-vision/internal/articles"
	"net/http"
)


func NewRouter(h *articles.Handler) {
	http.HandleFunc("/articles", h.Repo.CreateRepositoryHandler)
}