package usecase

import (
	"communication-server/infrastructure/postgresql/gen"
	"communication-server/internal/domain"
	"communication-server/internal/port"
	"context"

	"github.com/google/uuid"
)

type playerUsecase struct {
	repository port.Repository
}

type PlayerUsecase interface {
	FindPlayer(ctx context.Context, id uuid.UUID) (domain.Player, error)
	FindPlayerByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (domain.Player, error)
}

var _ PlayerUsecase = (*playerUsecase)(nil)

func NewPlayerUsecase(repository port.Repository) PlayerUsecase {
	return &playerUsecase{repository}
}

func (pu playerUsecase) FindPlayer(ctx context.Context, id uuid.UUID) (res domain.Player, err error) {
	var q gen.Querier
	ctx, q, err = pu.repository.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	res.Player, err = q.PlayerByID(ctx, id)
	if err != nil {
		err = pu.repository.RollbackTx(ctx, err)
		return
	}

	_, err = pu.repository.CommitTx(ctx)
	return
}

func (pu playerUsecase) FindPlayerByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (res domain.Player, err error) {
	res.Player, err = pu.repository.PlayerByEmailOrUsername(ctx, usernameOrEmail)
	return
}
