package repository

import (
	"context"
	"social-backend/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, post *domain.Post) error {

	query := `INSERT INTO posts(id, user_id, content, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, post.ID, post.UserID, post.Content, post.CreatedAt)
	return err
}

func (r *PostRepository) FindAll(ctx context.Context) ([]*domain.Post, error) {

	rows, err := r.db.Query(ctx, `SELECT id, user_id, content, created_at FROM posts ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostRepository) FindByID(ctx context.Context, id string) (*domain.Post, error) {
	var post domain.Post
	query := `SELECT id, user_id, content FROM posts WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&post.ID, &post.UserID, &post.Content)

	return &post, err
}

func (r *PostRepository) GetFeed(ctx context.Context) ([]domain.Post, error) {
	rows, err := r.db.Query(ctx, `SELECT id, user_id, content FROM posts ORDER BY id DESC LIMIT 10`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) Delete(ctx context.Context, id string, userID string) error {
	query := `DELETE FROM posts WHERE id = $1 AND user_id = $2`
	_, err := r.db.Exec(ctx, query, id, userID)
	return err
}
