package port

import (
	"communication-server/infrastructure/postgresql/gen"
	"context"
	"database/sql"
)

type Repository interface {
	gen.Querier

	// BeginTx begins a new transaction and puts it the returned context.
	//
	// If the tx is already created in the context, use it; otherwise,
	// begin a new one.
	BeginTx(ctx context.Context, opts *sql.TxOptions) (newCtx context.Context, q *gen.Queries, err error)

	// RollbackTx rollbacks the transaction with the given error.
	RollbackTx(ctx context.Context, err error) error

	// CommitTx commits the transaction.
	CommitTx(ctx context.Context) (newCtx context.Context, err error)
}
