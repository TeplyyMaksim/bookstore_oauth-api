package app

import (
	"github.com/TeplyyMaksim/bookstore_oauth-api/domain/access_token"
	"github.com/TeplyyMaksim/bookstore_oauth-api/http"
	"github.com/TeplyyMaksim/bookstore_oauth-api/repository/db"
	"github.com/labstack/echo"
)

var router = echo.New()

func StartApplication () {
	atRepository := db.NewRepository()
	atService := access_token.NewService(atRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Start(":8001")
}