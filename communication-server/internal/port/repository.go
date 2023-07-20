package port

import "communication-server/infrastructure/postgresql/gen"

type Repository interface {
	gen.Querier
}
