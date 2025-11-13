package articles

import (
	"context"
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertArticle(dto CreateArticleDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	insertQuery := `
		INSERT INTO posts (title, content, category, status)
		VALUES (?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, insertQuery, dto.Title, dto.Content, dto.Category, dto.Status)

	if err != nil {
		return err
	}

	return nil

}
