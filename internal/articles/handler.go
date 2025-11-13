package articles

import (
	"encoding/json"
	"hafiztri123/be-sharing-vision/internal/utils"
	"net/http"
)

type Handler struct {
	Repo *Repository
}

func NewHandler(r *Repository) *Handler {
	return &Handler{
		Repo: r,
	}
}

func (re *Repository) CreateRepositoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.NewJSONResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	var dto CreateArticleDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		utils.NewJSONResponse(w, "Bad Request", http.StatusBadRequest, nil)
		return
	}

	err = re.InsertArticle(dto)
	if err != nil {
		utils.NewJSONResponse(w, "Internal Server Error", http.StatusInternalServerError, nil)
	}

	utils.NewJSONResponse(w, "Success", http.StatusOK, nil)
}
