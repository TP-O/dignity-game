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
	authUsecase   usecase.AuthUsecase
	playerUsecase usecase.PlayerUsecase
}

func New(
	cfg config.App,
	cache port.Cache,
	mailer port.Mailer,
	authUsecase usecase.AuthUsecase,
	playerUsecase usecase.PlayerUsecase,
) *apiServer {
	if cfg.Env == config.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	return &apiServer{
		cfg,
		cache,
		authUsecase,
		playerUsecase,
	}
}

func (as apiServer) Use(router *gin.RouterGroup) {
	router.GET("/player/:id", as.getPlayerByID)
	router.POST("/auth/login", as.loginPlayer)
	router.POST("/auth/register", as.registerPlayer)
}
