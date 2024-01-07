package queryhandlers

import "context"

// IQueryHandler defines a generic interface for a query handler.
type IQueryHandler[TQuery any, TResult any] interface {
	Handle(ctx context.Context, query TQuery) (TResult, error)
}

// Other specific handler interfaces or methods can be defined as needed.
