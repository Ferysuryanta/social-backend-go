package service

import (
	"context"
	"errors"
	"social-backend/internal/domain"
	"social-backend/pkg/worker"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo   domain.UserRepository
	worker *worker.WorkerPool
}

func NewAuthService(r domain.UserRepository, w *worker.WorkerPool) *AuthService {
	return &AuthService{
		repo:   r,
		worker: w,
	}
}

func (s AuthService) Register(ctx context.Context, email string, password string) (*domain.User, error) {

	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return nil, errors.New("user exists")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	user := &domain.User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: string(hash),
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	_ = s.worker.Submit(func() {
		println("Send email async to: " + user.Email)
	})
	return user, err
}
