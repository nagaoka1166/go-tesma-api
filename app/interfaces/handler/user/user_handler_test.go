// app/interfaces/handler/user/user_handler_test.go

package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
    "github.com/nagaoka166/go-tesma-api/app/domain/entity"
	mock_repository "github.com/nagaoka166/go-tesma-api/app/domain/repository/mock"
	"github.com/nagaoka166/go-tesma-api/app/interfaces/handler/user"
)

func TestLoginHandler(t *testing.T) {
    // モックを用意
    mockCtrl := gomock.NewController(t)
    mockUsecase := mock_repository.NewMockUserUsecase(mockCtrl)

    // モックの振る舞いを定義
    mockUsecase.EXPECT().Login(gomock.Any(), "test@email.com", "password").Return(&entity.User{}, nil)

    // リクエストを模倣
    r := httptest.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(`{"email": "test@email.com", "password": "password"}`))
    w := httptest.NewRecorder()

    handler := handler.NewLoginHandler(mockUsecase)
    handler.Login(w, r)

    // レスポンスを確認
    if w.Code != http.StatusOK {
        t.Errorf("Expected HTTP status 200; got %d", w.Code)
    }
}


