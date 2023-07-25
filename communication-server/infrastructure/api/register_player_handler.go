package api

import (
	"communication-server/internal/dto"
	"communication-server/internal/presenter"

	"github.com/gin-gonic/gin"
)

func (as apiServer) RegisterPlayer(ctx *gin.Context) {
	var (
		req dto.RegisterPlayer
		res presenter.LoginPlayerPresenter
		err error
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	res, err = as.authUsecase.Register(ctx, req)
	if err != nil {
		return
	}

	go func() {
		if link, err := as.authUsecase.GenerateEmailVerificationLink(res.Player.ID); err == nil {
			as.mailer.SendEmailVerificationEmail(res.Player.Email, link)
		}
	}()

	ctx.JSON(200, map[string]any{
		"ok":   true,
		"data": res,
	})
}
