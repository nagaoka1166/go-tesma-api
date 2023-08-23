// :ファイル名:/app/interfaces/handler/user_handler.go
package handler

import (
	"net/http"
	"context"
	"github.com/nagaoka1166/go-tesma-api/app/domain/usecase"
	"github.com/nagaoka1166/go-tesma-api/app/domain/entity"
    // "strings"

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

func (h *UserHandler) SignUp(c *gin.Context) {
    var user entity.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
        return
    }

    idToken, err := h.UserUsecase.CreateUser(context.Background(), &user)
    if err != nil {
        switch err.Error() {
        case "user already exists":
            c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
        default:
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"status": "user created", "idToken": idToken})
}

func (h *UserHandler) Login (c *gin.Context) {
    var user entity.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    loggedInUser, idToken, err := h.UserUsecase.Login(context.Background(), user.Email, user.Password)
    if err != nil {
        switch err.Error() {
        case "user already exists":
            c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
        default:
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"status": "user logged in", "idtoken": idToken, "user": loggedInUser})
}