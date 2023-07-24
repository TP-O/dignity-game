package api

import (
	"communication-server/internal/dto"
	"communication-server/internal/presenter"

	"github.com/gin-gonic/gin"
)

func (as apiServer) LoginPlayer(ctx *gin.Context) {
	var (
		req dto.LoginPlayer
		res presenter.LoginPlayerPresenter
		err error
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	res, err = as.authUsecase.Login(ctx, req)
	if err != nil {
		return
	}

	ctx.JSON(200, map[string]any{
		"ok":   true,
		"data": res,
	})
}
