package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (as apiServer) getPlayerByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}

	player, err := as.playerUsecase.FindPlayer(ctx, id)
	if err != nil {
		return
	}

	ctx.JSON(200, map[string]any{
		"ok":   true,
		"data": player,
	})
}
