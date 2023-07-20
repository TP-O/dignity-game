package postgresql

import (
	"communication-server/config"
	"communication-server/infrastructure/postgresql/gen"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/avast/retry-go"
)

func New(cfg config.PostgreSQL) Store {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	))
	if err != nil {
		log.Panic(err)
	}

	if err := retry.Do(
		func() error {
			if err := db.Ping(); err != nil {
				log.Println(err.Error())
				return err
			}
			return nil
		},
		retry.Attempts(10),
		retry.DelayType(retry.RandomDelay),
		retry.MaxJitter(10*time.Second),
	); err != nil {
		log.Panic(err)
	}

	db.SetMaxOpenConns(cfg.PollSize)
	db.SetMaxIdleConns(cfg.PollSize)
	db.SetConnMaxIdleTime(0)

	return &store{
		Queries: gen.New(db),
		db:      db,
	}
}
