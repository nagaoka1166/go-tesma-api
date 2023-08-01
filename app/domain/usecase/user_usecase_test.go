
package usecase_test

import (
    "context"
    "testing"
    
    "github.com/golang/mock/gomock"
    
    "github.com/nagaoka166/go-tesma-api/app/domain/entity"
    "github.com/nagaoka166/go-tesma-api/app/domain/repository"
    "github.com/nagaoka166/go-tesma-api/app/domain/usecase" 
)

func TestCreateUser(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepo := repository.NewMockUserRepository(ctrl)

    user := &entity.User{
        Email:    "test@gmail.com",
        Password: "password",
    }

    mockUserRepo.EXPECT().GetUserByEmail(context.Background(), user.Email).Return(user, nil)

    usecase := usecase.NewUserUsecase(mockUserRepo)

    err := usecase.CreateUser(context.Background(), user)
    if err != nil {
        t.Fatalf("failed to create user: %v", err)
    }
}
