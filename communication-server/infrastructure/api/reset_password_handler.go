package api

import (
	"communication-server/internal/dto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (as apiServer) ResetPassword(ctx *gin.Context) {
	var (
		id        uuid.UUID
		expiredAt int
		req       dto.ResetPassword
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

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	err = as.authUsecase.ResetPassword(ctx, id, int64(expiredAt), ctx.Query("signature"), req)
	if err != nil {
		return
	}

	ctx.JSON(200, map[string]any{
		"ok":   true,
		"data": true,
	})
}
