package service

import (
	"context"
	"errors"
	"social-backend/internal/domain"
	"social-backend/internal/repository"
	"social-backend/pkg/jwt"
	"social-backend/pkg/worker"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo   repository.UserRepository
	worker *worker.WorkerPool
}

func NewAuthService(r repository.UserRepository, w *worker.WorkerPool) *AuthService {
	return &AuthService{
		repo:   r,
		worker: w,
	}
}

func (s AuthService) Register(ctx context.Context, email string, password string) (*domain.User, error) {

	existingUser, err := s.repo.FindByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, errors.New("user already exists")
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

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return jwt.Generate(user.ID)
}
