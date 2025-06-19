package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go-rest/internal/server"
)

func main() {
	cfg := &server.Config{
		Port:    9999,
		WasmDir: "./wasm",
	}

	srv := server.New(cfg)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv.Run(ctx)
}
