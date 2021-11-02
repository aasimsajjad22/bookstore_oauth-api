package app

import (
	"github.com/aasimsajjad22/bookstore_oauth-api/src/http"
	"github.com/aasimsajjad22/bookstore_oauth-api/src/repository/db"
	"github.com/aasimsajjad22/bookstore_oauth-api/src/repository/rest"
	"github.com/aasimsajjad22/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := access_token.NewService(rest.NewRestUsersRepository(), dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8081")
}
