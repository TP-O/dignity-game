package cache

import (
	"communication-server/internal/domain"
	"communication-server/internal/port"

	"github.com/dgraph-io/badger/v3"
)

type Cache struct {
	bg *badger.DB
}

func New(bg *badger.DB) port.Cache {
	return &Cache{
		bg,
	}
}

func (c Cache) AllPlayers() []domain.Player {
	return []domain.Player{}
}
