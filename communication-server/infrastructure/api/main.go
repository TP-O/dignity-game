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
	mailer        port.Mailer
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
		mailer,
		authUsecase,
		playerUsecase,
	}
}

func (as apiServer) Use(router *gin.RouterGroup) {
	router.GET("/player/:id", as.GetPlayerByID)
    router.GET("/player/:usernameOrEmail", as.GetPlayerByUsernameOrEmail)
	router.POST("/auth/login", as.LoginPlayer)
	router.POST("/auth/register", as.RegisterPlayer)

	router.POST("/auth/send_verification_email", as.SendEmailVerificationEmail)
	router.GET("/auth/verify", as.VerifyEmail)

	router.POST("/auth/send_reset_password_email", as.SendResetPasswordEmail)
	router.POST("/auth/reset_password", as.ResetPassword)
}
