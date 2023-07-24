package api

import (
	"communication-server/internal/dto"

	"github.com/gin-gonic/gin"
)

func (as apiServer) SendResetPasswordEmail(ctx *gin.Context) {
	var (
		req  dto.SendResetPasswordEmail
		link string
		err  error
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	_, err = as.playerUsecase.FindPlayerByUsernameOrEmail(ctx, req.Email)
	if err != nil {
		return
	}

	link, err = as.authUsecase.GenerateResetPasswordLink(ctx, req.Email)
	if err != nil {
		return
	}

	go as.mailer.SendResetPasswordEmail(req.Email, link)
	ctx.JSON(200, map[string]any{
		"ok":   true,
		"data": true,
	})
}
