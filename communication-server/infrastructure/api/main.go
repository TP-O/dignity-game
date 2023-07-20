package api

import (
	"communication-server/config"
	"communication-server/internal/port"
	"communication-server/internal/usecase"

	"github.com/gin-gonic/gin"
)

type apiServer struct {
	cfg           config.App
	cache         port.Cache
	playerUsecase usecase.PlayerUsecaseContract
}

func New(
	cfg config.App,
	cache port.Cache,
	playerUsecase usecase.PlayerUsecaseContract,
) *apiServer {
	if cfg.Env == config.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	return &apiServer{
		cfg,
		cache,
		playerUsecase,
	}
}

func (as apiServer) Use(router *gin.RouterGroup) {
	router.GET("/player/:id", as.getPlayerByID)
}
