package apis

import (
	"communication-server/internal/apps/models"
	"communication-server/internal/ports"
)

type Application struct {
	db ports.DbPort
}

func NewApplication(db ports.DbPort) *Application {
	return &Application{db: db}
}

func (a *Application) GetItems() ([]*models.Item, error) {
	items, err := a.db.GetItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (a *Application) AddItem(name string) error {
	err := a.db.AddItem(name)
	if err != nil {
		return err
	}
	return nil
} 