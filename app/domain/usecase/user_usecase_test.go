
package usecase_test

import (
    "context"
    "testing"
    
    "github.com/golang/mock/gomock"
    
    "app/domain/entity"
    "app/domain/entity/repository"
    "app/domain/usecase" 
)

func TestCreateUser(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepo := repository.NewMockUserRepository(ctrl)

    user := &entity.User{
        Email:    "test@gmail.com",
        Password: "password",
    }

    mockUserRepo.EXPECT().CreateUser(context.Background(), user).Return(nil)

    usecase := usecase.NewUserUsecase(mockUserRepo)

    err := usecase.CreateUser(context.Background(), user)
    if err != nil {
        t.Fatalf("failed to create user: %v", err)
    }
}
