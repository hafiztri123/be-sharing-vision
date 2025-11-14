package articles

import (
	"encoding/json"
	"hafiztri123/be-sharing-vision/internal/utils"
	"net/http"
	"strconv"
)

type Handler struct {
	Repo *Repository
}

func NewHandler(r *Repository) *Handler {
	return &Handler{
		Repo: r,
	}
}

func (h *Handler) CreateArticleHandler(w http.ResponseWriter, r *http.Request) {

	var dto CreateArticleDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		utils.NewJSONResponse(w, "Bad Request", http.StatusBadRequest, nil)
		return
	}

	validationErrors := dto.Validate()
	if validationErrors != nil {
		utils.NewJSONResponse(w, "Bad Request", http.StatusBadRequest, validationErrors)
		return
	}

	err = h.Repo.InsertArticle(dto)
	if err != nil {
		utils.NewJSONResponse(w, "Internal Server Error", http.StatusInternalServerError, nil)
		return
	}

	utils.NewJSONResponse(w, "Article created", http.StatusCreated, nil)
}

func (h *Handler) GetArticlesHandler(w http.ResponseWriter, r *http.Request) {

	pagiationParams, validationErrors := utils.ExtractPathValue(r)
	if validationErrors != nil {
		utils.NewJSONResponse(w, "Bad Request", http.StatusBadRequest, validationErrors)
		return
	}

	status := r.URL.Query().Get("status")

	result, err := h.Repo.GetArticlesPaginated(*pagiationParams, status)
	if err != nil {
		utils.NewJSONResponse(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	utils.NewJSONResponse(w, "Get Articles", http.StatusOK, result)
}

func (h *Handler) GetArticleHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.NewJSONResponse(
			w, "Bad Request", http.StatusBadRequest, nil,
		)
		return
	}

	result, err := h.Repo.GetArticle(id)
	if err != nil {
		if err.Error() == "Data not found" {
			utils.NewJSONResponse(w, "Data not found", http.StatusNotFound, nil)
			return
		}
		
		utils.NewJSONResponse(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	utils.NewJSONResponse(w, "Get Article", http.StatusOK, result)
}

func (h *Handler) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {

	var dto PutArticleDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		utils.NewJSONResponse(w, "Bad Request", http.StatusBadRequest, nil)
		return
	}

	if validationErrors := dto.Validate(); validationErrors != nil {
		utils.NewJSONResponse(w, "Bad Request", http.StatusBadRequest, validationErrors)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.NewJSONResponse(w, "Id not valid", http.StatusBadRequest, nil)
		return
	}

	err = h.Repo.UpdateArticle(id, dto)
	if err != nil {
		if err.Error() == "Data not found" {
			utils.NewJSONResponse(w, "Data not found", http.StatusNotFound, nil)
			return
		}

		utils.NewJSONResponse(w, "Internal Server Error", http.StatusInternalServerError, nil)
		return
	}

	utils.NewJSONResponse(w, "Article updated", http.StatusOK, nil)
}

func (h *Handler) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.NewJSONResponse(w, "Invalid Id", http.StatusBadRequest, nil)
		return
	}

	err = h.Repo.DeleteArticle(id)
	if err != nil {
		if err.Error() == "Data not found" {
			utils.NewJSONResponse(w, "Data not found", http.StatusNotFound, nil)
			return
		}

		utils.NewJSONResponse(w, "Internal Server Error", http.StatusInternalServerError, nil)
		return
	}

	utils.NewJSONResponse(w, "Article Deleted", http.StatusNoContent, nil)
}
