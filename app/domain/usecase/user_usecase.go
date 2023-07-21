package usecase

import (
    "context"
    
    "app/domain/entity"
    "app/domain/repository"
)

type UserUsecase interface {
    CreateUser(ctx context.Context, user *entity.User) error
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
    if err == nil {
        return errors.New("user already exists")
    }

    // Create the user
    err = u.userRepo.CreateUser(ctx, user)
    if err != nil {
        return err
    }

    return nil
}
