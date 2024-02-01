package repository

import (
	"context"
	"lrucache/domain"
)

type UserRepository interface {
	InsertOrUpdateUser(ctx context.Context, user domain.User) error
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
}
