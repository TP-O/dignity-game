package postgresql

import (
	"communication-server/config"
	"communication-server/infrastructure/postgresql/gen"
	"communication-server/pkg"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/avast/retry-go"
)

type store struct {
	*gen.Queries
	db *sql.DB
}

func New(cfg config.PostgreSQL) *store {
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

// Close closes database connection.
func (s *store) Close() error {
	return s.db.Close()
}

// BeginTx begins a new transaction and puts it the returned context.
//
// If the transaction is already created in the context, use it; otherwise,
// begin a new one.
func (s *store) BeginTx(ctx context.Context, opts *sql.TxOptions) (newCtx context.Context, q *gen.Queries, err error) {
	var (
		c  int
		ok bool
		tx *sql.Tx
	)

	if tx, ok = ctx.Value(pkg.TxContextKey).(*sql.Tx); ok {
		if c, ok = ctx.Value(pkg.TxContextCallKey).(int); !ok {
			return ctx, nil, pkg.ErrUnknownTxCounter
		}

		newCtx = context.WithValue(ctx, pkg.TxContextCallKey, c+1)
		if q, ok = ctx.Value(pkg.QuerierContextKey).(*gen.Queries); ok {
			return newCtx, q, nil
		}

		q = gen.New(tx)
		newCtx = context.WithValue(newCtx, pkg.QuerierContextKey, q)
		return
	}

	tx, err = s.db.BeginTx(ctx, opts)
	if err != nil {
		return
	}

	q = gen.New(tx)
	newCtx = context.WithValue(ctx, pkg.QuerierContextKey, q)
	newCtx = context.WithValue(newCtx, pkg.TxContextKey, tx)
	newCtx = context.WithValue(newCtx, pkg.TxContextCallKey, 0)
	return
}

// RollbackTx rollbacks the transaction in the given context
// with the given error.
func (s *store) RollbackTx(ctx context.Context, err error) error {
	var (
		tx *sql.Tx
		ok bool
	)

	if tx, ok = ctx.Value(pkg.TxContextKey).(*sql.Tx); !ok {
		return pkg.ErrUnknownTx
	}

	if rbErr := tx.Rollback(); rbErr != nil {
		return rbErr
	}

	return err
}

// CommitTx commits the transaction in the given context.
//
// If the transaction is begun many times, only beginner can commit it.
// Others just reduce the call counter.
//
// NOTE: make sure that this is called at the end of function using BeginTx.
func (s *store) CommitTx(ctx context.Context) (newCtx context.Context, err error) {
	var (
		tx *sql.Tx
		c  int
		ok bool
	)

	if tx, ok = ctx.Value(pkg.TxContextKey).(*sql.Tx); !ok {
		return ctx, pkg.ErrUnknownTx
	}

	if c, ok = ctx.Value(pkg.TxContextCallKey).(int); !ok {
		return ctx, pkg.ErrUnknownTxCounter
	}

	if c == 0 {
		return ctx, tx.Commit()
	}

	newCtx = context.WithValue(ctx, pkg.TxContextCallKey, c-1)
	return
}
