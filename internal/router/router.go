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

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})

}
