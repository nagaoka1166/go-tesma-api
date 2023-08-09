// : github.com/nagaoka166/go-tesma-api/app/interfaces/handler/user_handler.go
package handler

import (
	"log"
	"net/http"
	"context"
	"github.com/nagaoka166/go-tesma-api/app/domain/usecase"
	"github.com/nagaoka166/go-tesma-api/app/domain/entity"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: u,
	}
}
// "error": "user already exists"

func (h *UserHandler) SignUp(c *gin.Context) {
    var user entity.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
        return
    }

    exists, err := h.UserUsecase.UserExists(context.Background(), user.Email)
if err != nil {
    log.Printf("error in UserExists: %v", err)
    // ここでFirebaseから返されるエラーメッセージをそのままレスポンスに含めるように変更
    c.JSON(http.StatusInternalServerError, gin.H{"error": "UserExists error: " + err.Error()})
    return
}
    if exists {
        c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
        return
    }

    err = h.UserUsecase.CreateUser(context.Background(), &user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"status": "user created"})
}
