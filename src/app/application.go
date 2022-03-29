package app

import (
	"test3/hariprathap-hp/bookstore_oauth_go/src/http_handler"
	"test3/hariprathap-hp/bookstore_oauth_go/src/repository/db"
	"test3/hariprathap-hp/bookstore_oauth_go/src/token_service"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := token_service.NewService(db.NewRepository())
	atHandler := http_handler.NewHTTPHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetbyID)
	router.Run(":8081")
}
