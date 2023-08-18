package usecase_test

import (
    "context"
    "testing"
    "errors"
    "github.com/golang/mock/gomock"
    
    "github.com/nagaoka166/go-tesma-api/app/domain/entity"
    "github.com/nagaoka166/go-tesma-api/app/domain/usecase" 
    mock_repository "github.com/nagaoka166/go-tesma-api/app/domain/repository/mock"
)

func TestCreateUser(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepo := mock_repository.NewMockUserRepository(ctrl)

    user := &entity.User{
        Email:    "111111@ed.ritsumei.ac.jp",
        Password: "password1234",
    }

    // GetUserByEmailのmockを設定。ユーザーが存在しないと仮定して、エラーを返す。
    // mockUserRepo.EXPECT().UserExists(context.Background(), user.Email).Return(nil, errors.New("user not found")).Times(1)

    // CreateUserのmockを設定
    mockUserRepo.EXPECT().CreateUser(context.Background(), user).Return(nil).Times(1)
    mockUserRepo.EXPECT().UserExists(context.Background(), user.Email).Return(false, nil).Times(1)

    usecase := usecase.NewUserUsecase(mockUserRepo)

    err := usecase.CreateUser(context.Background(), user)
    if err != nil {
        t.Fatalf("failed to create user: %v", err)
    }
}


