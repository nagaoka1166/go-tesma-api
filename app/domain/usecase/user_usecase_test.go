// app/domain/usecase/user_usecase_test.go
package usecase_test

import (
	"context"
    "testing"
    "strings"

	"github.com/golang/mock/gomock"
	"github.com/nagaoka1166/go-tesma-api/app/domain/entity"
	"github.com/nagaoka1166/go-tesma-api/app/domain/usecase"
	mock_repository "github.com/nagaoka1166/go-tesma-api/app/domain/repository/mock"
)

func FuzzCreateUser(f *testing.F) {
    ctrl := gomock.NewController(f)
    defer ctrl.Finish()

    mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
    ctx := context.Background()
    userUsecase := usecase.NewUserUsecase(mockUserRepo)

    // f.Add() の使用は問題がない場合にのみコメントアウトを解除
    // f.Add("111111@ed.ritsumei.ac.jp", "pass12345")

    f.Fuzz(func(t *testing.T, email, password string) {
        user := &entity.User{
            Email:    email,
            Password: password,
        }

        mockUserRepo.EXPECT().CreateUser(ctx, user).Return(nil).AnyTimes()
        mockUserRepo.EXPECT().UserExists(ctx, user.Email).Return(false, nil).AnyTimes()

        idToken, err := userUsecase.CreateUser(ctx, user)
        if err != nil {
            t.Fatalf("failed to create user: %v", err)
        }
        _ = idToken 
        
    })
}


// MEMO: VerifyIDTokenメソッドはおそらく必要ないが一旦残しておく
// func FuzzVerifyIDToken(f *testing.F) {
// 	ctrl := gomock.NewController(f)
// 	defer ctrl.Finish()

	
// 	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
// 	ctx := context.Background()

// 	dummyToken := &auth.Token{
// 		UID: "dummyUID",
// 	}

// 	f.Fuzz(func(t *testing.T, testIDToken string) {
// 		...
// 		token, err := mockUserRepo.VerifyIDToken(ctx, testIDToken)
// 		if token != nil && contains(testIDToken, "valid") && (err != nil || token.UID != dummyToken.UID) {
// 			t.Fatalf("expected UID %v with no error, got UID %v with error %v", dummyToken.UID, token.UID, err)
// 		} else if !contains(testIDToken, "valid") && err == nil {
// 			t.Fatalf("expected an error for invalid token but got UID %v", token.UID)
// 		}
// 	})
	
// }

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}