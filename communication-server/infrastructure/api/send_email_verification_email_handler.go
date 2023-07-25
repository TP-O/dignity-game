package api

import (
	"communication-server/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (as apiServer) SendEmailVerificationEmail(ctx *gin.Context) {
	var (
		player domain.Player
		link   string
		err    error
	)

	player, err = as.playerUsecase.FindPlayer(ctx, uuid.MustParse(ctx.GetString("id")))
	if err != nil {
		return
	}

	link, err = as.authUsecase.GenerateEmailVerificationLink(player.ID)
	if err != nil {
		return
	}

	go as.mailer.SendEmailVerificationEmail(player.Email, link)
	ctx.JSON(200, map[string]any{
		"ok":   true,
		"data": true,
	})
}
