package api

import (
	"communication-server/internal/domain"
	"github.com/gin-gonic/gin"
)

func (as apiServer) GetPlayerByUsernameOrEmail(ctx *gin.Context) {
	var (
		usernameOrEmail string
		res domain.Player
		err error
	)

	usernameOrEmail = ctx.Param("usernameOrEmail")

	res, err = as.playerUsecase.FindPlayerByUsernameOrEmail(ctx, usernameOrEmail)
	if err != nil {
		return
	}

	ctx.JSON(200, map[string]any{
		"ok":   true,
		"data": res,
	})
}
