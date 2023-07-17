package gingonic

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"chat-server/internal/ports"
	"chat-server/configs"
)

type Adapter struct {
	conf *configs.Configs
	api ports.ApiPort
}

func NewAdapter(conf *configs.Configs, a ports.ApiPort) *Adapter {
	return &Adapter {
		conf: conf,
		api: a,
	}
}

func (a *Adapter) Run() {
	server := gin.Default()
	server.Group("/api/chat")

	server.NoRoute(func (ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "fail",
			"message": fmt.Sprintf("Route %s not found", ctx.Request.URL),
		})
	})

	server.POST("/items", a.addItem)

	log.Fatal(server.Run(":" + a.conf.Port))
}