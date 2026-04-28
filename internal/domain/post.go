package domain

import "time"

type Post struct {
	ID        string
	UserID    string
	Content   string
	CreatedAt time.Time
}
