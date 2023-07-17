package sqlc

import (
	"log"
	"database/sql"
	"chat-server/configs"
	"chat-server/internal/apps/models"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(conf *configs.Configs) *Database {
	db, err := sql.Open(conf.PgDriverName, conf.PgDataSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err.Error())
	}

	return &Database{
		db: db,
	}
}

func (db *Database) AddItem(name string) (err error) {
	_, err = db.db.Exec("INSERT INTO items (name) VALUES ($1)", name)
	return
}

func (db *Database) GetItems() (items []*models.Item, err error) {
	rows, err := db.db.Query("SELECT name FROM items")
	if err != nil {
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return
		}

		items = append(items, &models.Item{
			Name: name,
		})
	}

	return
}