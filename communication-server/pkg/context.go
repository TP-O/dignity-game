package pkg

type ContextKeyType string

const (
	TxContextKey      = ContextKeyType("tx_context")
	TxContextCallKey  = ContextKeyType("tx_context_call")
	QuerierContextKey = ContextKeyType("querier_context")
)
