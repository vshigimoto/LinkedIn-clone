package repository

import (
	"context"
	entity "github.com/vshigimoto/Blog/internal/blog/entity/user"
)

type Repository struct {
	UserRepo
}

type UserRepo interface {
	Create(ctx context.Context, user entity.User) (int, error)
	Get(ctx context.Context, userID int) (entity.User, error)
	Delete(ctx context.Context, userID int) error
	Update(ctx context.Context, user entity.User) error
}
