package ports

import "communication-server/internal/apps/models"

type ApiPort interface {
	AddItem(name string) error
	GetItems() ([]*models.Item, error)
}