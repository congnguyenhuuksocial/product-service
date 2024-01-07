package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"product-service/cmd/container"
	"syscall"
)

func main() {
	bootstarp()
}

func bootstarp() {
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, os.Interrupt)
	app := fx.New(
		container.Modules,
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)
	if err := app.Start(ctx); err != nil {
		fmt.Println(err)
	}
	<-app.Done()
	defer stop()
}
