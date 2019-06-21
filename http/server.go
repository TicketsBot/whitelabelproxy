package http

import (
	"github.com/TicketsBot/whitelabelproxy/config"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.POST("/", RequestHandler)

	if err := router.Run(config.Conf.Server.Host); err != nil {
		panic(err)
	}
}
