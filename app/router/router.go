package router

import (
	"net/http"
	"github.com/nagaoka1166/go-tesma-api/app/domain/usecase"
	"github.com/nagaoka1166/go-tesma-api/app/interfaces/handler/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(u usecase.UserUsecase) *gin.Engine {
	r := gin.Default()
    
	v1 := r.Group("/v1")
	{
	userHandler := handler.NewUserHandler(u)
	v1.POST("/signup", userHandler.SignUp)
	v1.POST("/login", userHandler.Login)


	v1.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	}

	return r
}
