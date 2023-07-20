package usecase

import (
	"communication-server/internal/domain"
	"communication-server/internal/port"
	"context"
)

type playerUsecase struct {
	repository port.Repository
}

type PlayerUsecaseContract interface {
	FindPlayer(ctx context.Context, id int) (domain.Player, error)
}

var _ PlayerUsecaseContract = (*playerUsecase)(nil)

func NewPlayerUsecase(repository port.Repository) PlayerUsecaseContract {
	return &playerUsecase{repository}
}

func (pu playerUsecase) FindPlayer(ctx context.Context, id int) (domain.Player, error) {
	var data domain.Player
	res, err := pu.repository.PlayerByID(ctx, int64(id))
	data.ID = res.ID
	data.Username = res.Username.String

	return data, err
}
