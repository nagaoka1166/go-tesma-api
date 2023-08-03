package main

import (
	// "net/http
	"log"
	"github.com/nagaoka166/go-tesma-api/app/domain/usecase"
	infrastructure "github.com/nagaoka166/go-tesma-api/app/infrastructure/repository"
	"github.com/nagaoka166/go-tesma-api/app/router"
    // "github.com/gin-gonic/gin"

)


func main() {
	log.Println("start server...")

	userRepo := infrastructure.NewUserRepo() // ここは具体的なリポジトリの初期化に変えてください
	userUsecase := usecase.NewUserUsecase(userRepo)

	// r.POST("/test", func(c *gin.Context) {
	// 	// ここに新たなPOSTメソッドの処理を書くことができます。
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Test POST method successful!",
	// 	})
	// })
	

	// ルーターの初期化とサーバーの起動
	r := router.NewRouter(userUsecase)
	log.Fatal(r.Run())
}
