package service

import (
	"context"
	"social-backend/internal/domain"
	"social-backend/internal/repository"

	"github.com/google/uuid"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(r *repository.PostRepository) *PostService {
	return &PostService{repo: r}
}

func (s *PostService) Create(ctx context.Context, userID, content string) error {
	post := &domain.Post{
		ID:      uuid.New().String(),
		UserID:  userID,
		Content: content,
	}
	return s.repo.Create(ctx, post)
}

func (s *PostService) GetById(ctx context.Context, id string) (*domain.Post, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *PostService) Feed(ctx context.Context) ([]domain.Post, error) {
	return s.repo.GetFeed(ctx)
}
