package api

import (
	"communication-server/internal/dto"
	"communication-server/internal/presenter"

	"github.com/gin-gonic/gin"
)

func (as apiServer) registerPlayer(ctx *gin.Context) {
	var (
		req dto.RegisterPlayerDto
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

	ctx.JSON(200, map[string]any{
		"ok":   true,
		"data": res,
	})
}
