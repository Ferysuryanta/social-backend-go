package repository

import (
	"context"
	"social-backend/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query,
		user.ID, user.Email, user.Password)
	return err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`

	var user domain.User
	err := r.db.QueryRow(context.Background(), query, email).
		Scan(&user.ID, &user.Email, &user.Password)

	return &user, err

}
