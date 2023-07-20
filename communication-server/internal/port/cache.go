package port

import "communication-server/internal/domain"

type Cache interface {
	AllPlayers() []domain.Player
}
