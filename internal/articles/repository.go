package articles

import (
	"context"
	"database/sql"
	"errors"
	"hafiztri123/be-sharing-vision/internal/utils"
	"math"
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

func (r *Repository) GetArticlesPaginated(queryParams utils.PaginationParams, status string) (*utils.PaginationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	articles := []Article{}

	offset := ((queryParams.Offset - 1) * queryParams.Limit)

	fetchQuery := `
		SELECT id, title, content, category, status FROM 
		posts WHERE status = ? LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, fetchQuery, status, queryParams.Limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var article Article

		err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.Category, &article.Status)

		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	countQuery := `
		SELECT COUNT(*) FROM posts
		WHERE status = ?

	`

	var totalCount int

	err = r.db.QueryRowContext(ctx, countQuery, status).Scan(&totalCount)

	totalPages := 0
	if totalCount > 0 {
		totalPages = int(math.Ceil(float64(totalCount) / float64(queryParams.Limit)))
	}

	return &utils.PaginationResponse{
		Data: articles,
		PaginationMeta: utils.PaginationMeta{
			TotalRecords: totalCount,
			TotalPages:   totalPages,
		},
	}, nil
}

func (r *Repository) GetArticle(id int) (*Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fetchQuery := `
		SELECT id, title, content, category, status
		FROM posts WHERE id = ? 
	`

	var article Article

	err := r.db.QueryRowContext(ctx, fetchQuery, id).Scan(
		&article.Id,
		&article.Title,
		&article.Content,
		&article.Category,
		&article.Status,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Data not found")
		}
		return nil, err
	}

	return &article, nil
}

func (r *Repository) UpdateArticle(id int, dto PutArticleDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateQuery := `
		UPDATE posts
		SET title = ?, content = ?, category = ?, status = ?
		WHERE id = ?
	`

	result, err := r.db.ExecContext(ctx, updateQuery, dto.Title, dto.Content, dto.Category, dto.Status, id)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Data not found")
	}

	return nil
}

func (r *Repository) DeleteArticle(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	deleteQuery := `
		DELETE FROM posts
		WHERE id = ?
	`

	row, err := r.db.ExecContext(ctx, deleteQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Data not found")
	}

	return nil
}
