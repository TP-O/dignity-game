package ports

import "chat-server/internal/apps/models"

type ApiPort interface {
	AddItem(name string) error
	GetItems() ([]*models.Item, error)
}