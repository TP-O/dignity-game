package usecase

import (
	"communication-server/config"
	"communication-server/infrastructure/postgresql/gen"
	"communication-server/internal/domain"
	"communication-server/internal/dto"
	"communication-server/internal/port"
	"communication-server/internal/presenter"
	"communication-server/pkg"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/bcrypt"
)

const (
	TOKEN_LIFETIME = 24 * time.Hour
	BCRYPT_COST    = 20
)

type authUsecase struct {
	appCfg     config.App
	repository port.Repository
}

type AuthUsecase interface {
	Login(ctx context.Context, data dto.LoginPlayerDto) (presenter.LoginPlayerPresenter, error)
	Register(ctx context.Context, data dto.RegisterPlayerDto) (presenter.LoginPlayerPresenter, error)
	GenerateEmailVerificationLink(id uuid.UUID) (string, error)
	VerifyEmail(ctx context.Context, id uuid.UUID, expiredAt int64, signature string) error
}

var _ AuthUsecase = (*authUsecase)(nil)

func NewAuthentiUsecase(appCfg config.App, repository port.Repository) AuthUsecase {
	return &authUsecase{appCfg, repository}
}

func (au authUsecase) Login(ctx context.Context, data dto.LoginPlayerDto) (res presenter.LoginPlayerPresenter, err error) {
	var (
		player    domain.Player
		jsonToken paseto.JSONToken
	)

	if data.Email != "" {
		player.Player, err = au.repository.PlayerByEmailOrUsername(ctx, data.Email)
	} else {
		player.Player, err = au.repository.PlayerByEmailOrUsername(ctx, data.Username)
	}
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(data.Password)); err != nil {
		return
	}

	now := time.Now()
	exp := now.Add(TOKEN_LIFETIME)
	jsonToken = paseto.JSONToken{
		Issuer: "login",
		// Jti:        player.ID.String(), // TODO: implement jti checking
		Subject:    player.ID.String(),
		IssuedAt:   now,
		Expiration: exp,
	}

	if res.Token, err = paseto.NewV2().Encrypt([]byte(au.appCfg.SecretKey), jsonToken, nil); err != nil {
		return
	}

	res.Player = player
	return
}

func (au authUsecase) Register(ctx context.Context, data dto.RegisterPlayerDto) (res presenter.LoginPlayerPresenter, err error) {
	var (
		player         domain.Player
		hashedPassword []byte
		jsonToken      paseto.JSONToken
	)

	if _, err = au.repository.PlayerByEmailOrUsername(ctx, data.Email); err != sql.ErrNoRows {
		if err == nil {
			err = pkg.ErrEmailUnavailable
		}
		return
	}

	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(data.Password), BCRYPT_COST); err != nil {
		return
	}

	if player.Player, err = au.repository.CreatePlayer(ctx, gen.CreatePlayerParams{
		Username: pkg.RandomString(12),
		Email:    data.Email,
		Password: string(hashedPassword),
	}); err != nil {
		return
	}

	now := time.Now()
	exp := now.Add(TOKEN_LIFETIME)
	jsonToken = paseto.JSONToken{
		Issuer: "register",
		// Jti:        player.ID.String(), // TODO: implement jti checking
		Subject:    player.ID.String(),
		IssuedAt:   now,
		Expiration: exp,
	}

	if res.Token, err = paseto.NewV2().Encrypt([]byte(au.appCfg.SecretKey), jsonToken, nil); err != nil {
		return
	}

	res.Player = player
	return
}

func (au authUsecase) GenerateEmailVerificationLink(id uuid.UUID) (link string, err error) {
	var (
		expiredAt int64
		signature string
	)

	expiredAt = time.Now().Add(1 * time.Hour).UnixMilli()
	signature, err = pkg.SignWithHMAC(
		au.appCfg.SecretKey,
		fmt.Sprintf(`{"id":"%s","expiredAt":%d}`, id.String(), expiredAt),
	)
	if err != nil {
		return
	}

	link = fmt.Sprintf("%s/auth/verify?id=%s&expiredAt=%d&signature=%s",
		au.appCfg.Host,
		id.String(),
		expiredAt,
		signature,
	)
	return
}

func (au authUsecase) VerifyEmail(ctx context.Context, id uuid.UUID, expiredAt int64, signature string) (err error) {
	var (
		expectedSignature string
	)

	expectedSignature, err = pkg.SignWithHMAC(
		au.appCfg.SecretKey,
		fmt.Sprintf(`{"id":"%s","expiredAt":%d}`, id.String(), expiredAt),
	)
	if err != nil {
		return
	}

	if expectedSignature != signature {
		return pkg.ErrInvalidSignature
	}

	if time.Now().UnixMilli() > expiredAt {
		return pkg.ErrExpiredVersion
	}

	return au.repository.VerifyPlayerEmail(ctx, id)
}
