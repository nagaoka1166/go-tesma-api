package router

import (
	"app/domain/usecase"
	"app/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(u usecase.UserUsecase) *gin.Engine {
	r := gin.Default()

	userHandler := handler.NewUserHandler(u)
	r.POST("/signup", userHandler.SignUp)

	return r
}
