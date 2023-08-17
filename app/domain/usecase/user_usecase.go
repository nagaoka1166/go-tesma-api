package usecase

import (
    "log"
    "context"
    "errors"

    "github.com/nagaoka166/go-tesma-api/app/domain/entity"
    "github.com/nagaoka166/go-tesma-api/app/domain/repository"
    // mock_repository "github.com/nagaoka166/go-tesma-api/app/domain/repository/mock"
)

type UserUsecase interface {
    CreateUser(ctx context.Context, user *entity.User) error
    UserExists(ctx context.Context, email string) (bool, error)
}

type userUsecase struct {
    userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
    return &userUsecase{
        userRepo: userRepo,
    }
}

func (u *userUsecase) CreateUser(ctx context.Context, user *entity.User) error {
    // Validate the user details
    err := user.Validate()
    if err != nil {
        return err
    }

    // Check if the user already exists
    _, err = u.userRepo.GetUserByEmail(ctx, user.Email)
    if err != nil {
        // もしエラーが「ユーザーが見つからない」というエラーではない場合、エラーを返す
        if err.Error() != "user not found" {
            return err
        }
    } else {
        // ユーザーが見つかった場合、エラーを返す
        return errors.New("user already exists")
    }

    // Create the user
    err = u.userRepo.CreateUser(ctx, user)
    if err != nil {
        return err
    }

    return nil
}


func (u *userUsecase) UserExists(ctx context.Context, email string) (bool, error) {
    exists, err := u.userRepo.UserExists(ctx, email)
    if err != nil {
        log.Printf("error in userRepo.UserExists: %v", err)
    }
    return exists, err
}