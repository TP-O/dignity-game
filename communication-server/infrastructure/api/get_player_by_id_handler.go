package api

import (
	"communication-server/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (as apiServer) getPlayerByID(ctx *gin.Context) {
	var (
		id  uuid.UUID
		res domain.Player
		err error
	)

	id, err = uuid.Parse(ctx.Param("id"))
	if err != nil {
		return
	}

	res, err = as.playerUsecase.FindPlayer(ctx, id)
	if err != nil {
		return
	}

	ctx.JSON(200, map[string]any{
		"ok":   true,
		"data": res,
	})
}
