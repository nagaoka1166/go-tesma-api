// /go-tesma-api/app/domain/repository/user_repository.go
package repository

import (
	"github.com/nagaoka166/go-tesma-api/app/domain/entity"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
	Login(ctx context.Context, email, password string) (*entity.User, error)
	UserExists(ctx context.Context, email string) (bool, error)
}
