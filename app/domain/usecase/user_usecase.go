// ファイル名: app/domain/usecaseuser_usecase.go
package usecase

import (
    "log"
    "context"
    "errors"

    "github.com/nagaoka1166/go-tesma-api/app/domain/entity"
    "github.com/nagaoka1166/go-tesma-api/app/domain/repository"
    // mock_repository "github.com/nagaoka166/go-tesma-api/app/domain/repository/mock"
)

type UserUsecase interface {
    CreateUser(ctx context.Context, user *entity.User) (string, error)
    UserExists(ctx context.Context, email string) (bool, error)
    // RefreshUserToken(ctx context.Context, refreshToken string) (string, error)
    Login(ctx context.Context, email, password string) (*entity.User, string, error)
    
}

type userUsecase struct {
    userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
    return &userUsecase{
        userRepo: userRepo,
    }
}

func (u *userUsecase) CreateUser(ctx context.Context, user *entity.User) (string, error) {
    err := user.Validate()
    if err != nil {
        return "", err
    }
    exists, err := u.userRepo.UserExists(ctx, user.Email)
    if err != nil {
        return "",err
    }
    if exists {
        return "", errors.New("user already exists")
    }

    idToken, err := u.userRepo.CreateUser(ctx, user)
    if err != nil {
        return "", err
    }

    return idToken, nil
}

func (u *userUsecase) RefreshUserToken(ctx context.Context, refreshToken string) (string, error) {
    newToken, err := u.userRepo.RefreshToken(ctx, refreshToken)
    if err != nil {
        return "", err
    }
    return newToken, nil
}


func (u *userUsecase) UserExists(ctx context.Context, email string) (bool, error) {
    exists, err := u.userRepo.UserExists(ctx, email)
    if err != nil {
        log.Printf("error in userRepo.UserExists: %v", err)
    }
    return exists, err
}

func (u *userUsecase) Login(ctx context.Context, email, password string) (*entity.User, string, error) {
    user, idToken, err := u.userRepo.Login(ctx, email, password)
    if err != nil {
        return nil, "", errors.New("invalid login")
    }
    return user, idToken, nil
}