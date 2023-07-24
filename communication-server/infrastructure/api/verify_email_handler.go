package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (as apiServer) VerifyEmail(ctx *gin.Context) {
	var (
		id        uuid.UUID
		expiredAt int
		err       error
	)

	id, err = uuid.Parse(ctx.Query("id"))
	if err != nil {
		return
	}

	expiredAt, err = strconv.Atoi(ctx.Query("expiredAt"))
	if err != nil {
		return
	}

	err = as.authUsecase.VerifyEmail(ctx, id, int64(expiredAt), ctx.Query("signature"))
	if err != nil {
		return
	}

	ctx.JSON(200, map[string]any{
		"ok":   true,
		"data": true,
	})
}
