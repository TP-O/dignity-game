package postgresql

import (
	"communication-server/infrastructure/postgresql/gen"
	"context"
	"database/sql"
)

type Store interface {
	gen.Querier

	Close() error
}

type store struct {
	*gen.Queries
	db *sql.DB
}

func (s *store) Close() error {
	return s.db.Close()
}

func (s *store) execTx(ctx context.Context, fn func(q *gen.Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := gen.New(tx)
	if err = fn(q); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
