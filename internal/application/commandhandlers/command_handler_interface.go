package commandhandlers

import "context"

type ICommandHandler[T any] interface {
	Handle(ctx context.Context, command T) error
}
