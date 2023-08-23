package main

import (
	"log"
	"github.com/nagaoka166/go-tesma-api/app/domain/usecase"
	infrastructureDb "github.com/nagaoka166/go-tesma-api/app/infrastructure"
	infrastructure "github.com/nagaoka166/go-tesma-api/app/infrastructure/repository/user"
	"github.com/nagaoka166/go-tesma-api/app/router"
    // "github.com/gin-gonic/gin"

)


func main() {
	log.Println("start server...")
	dbInstance := infrastructureDb.InitDB()

	userRepo := infrastructure.NewUserRepo(dbInstance)
	userUsecase := usecase.NewUserUsecase(userRepo)

	

	// ルーターの初期化とサーバーの起動
	r := router.NewRouter(userUsecase)
	log.Fatal(r.Run())
}
