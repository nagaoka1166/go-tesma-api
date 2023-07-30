package main

import (
	"log"
	"app/domain/usecase"
	"app/router"
	"app/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("start server...")

	userRepo := infrastructure.NewUserRepo() // ここは具体的なリポジトリの初期化に変えてください
	userUsecase := usecase.NewUserUsecase(userRepo)

	// ルーターの初期化とサーバーの起動
	r := router.NewRouter(userUsecase)
	log.Fatal(r.Run())
}
