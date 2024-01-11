package commandhandlers

import "context"

type ICommandHandler[T any, TResult any] interface {
	Handle(ctx context.Context, command T) (TResult, error)
}
