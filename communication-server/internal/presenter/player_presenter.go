package presenter

import "communication-server/internal/domain"

type LoginPlayerPresenter struct {
	Token  string        `json:"token"`
	Player domain.Player `json:"player"`
}
