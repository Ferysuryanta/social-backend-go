package service

import (
	"context"
	"social-backend/internal/domain"
	"social-backend/internal/repository"
	"time"

	"github.com/google/uuid"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(r *repository.PostRepository) *PostService {
	return &PostService{repo: r}
}

func (s *PostService) Create(ctx context.Context, userID, content string) (*domain.Post, error) {
	post := &domain.Post{
		ID:        uuid.New().String(),
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}

	err := s.repo.Create(ctx, post)
	return post, err
}

func (s *PostService) GetAll(ctx context.Context) ([]*domain.Post, error) {
	return s.repo.FindAll(ctx)
}

func (s *PostService) GetById(ctx context.Context, id string) (*domain.Post, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *PostService) Feed(ctx context.Context) ([]domain.Post, error) {
	return s.repo.GetFeed(ctx)
}

func (s *PostService) Delete(ctx context.Context, id string, userID string) error {
	return s.repo.Delete(ctx, id, userID)
}
