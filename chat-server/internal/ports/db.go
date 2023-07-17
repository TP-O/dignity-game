package ports

import "chat-server/internal/apps/models"

type DbPort interface {
	AddItem(name string) error
	GetItems() ([]*models.Item, error)
}