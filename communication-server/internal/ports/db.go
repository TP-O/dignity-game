package ports

import "communication-server/internal/apps/models"

type DbPort interface {
	AddItem(name string) error
	GetItems() ([]*models.Item, error)
}