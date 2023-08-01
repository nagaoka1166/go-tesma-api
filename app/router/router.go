package router

import (
	"net/http"
	"github.com/nagaoka166/go-tesma-api/app/domain/usecase"
	"github.com/nagaoka166/go-tesma-api/app/interfaces/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(u usecase.UserUsecase) *gin.Engine {
	r := gin.Default()

	userHandler := handler.NewUserHandler(u)
	r.POST("/signup", userHandler.SignUp)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	return r
}
