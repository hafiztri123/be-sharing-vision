package articles



type CreateArticleDTO struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
	Status string `json:"status"`
}

type PutArticleDTO = CreateArticleDTO
type Article = CreateArticleDTO