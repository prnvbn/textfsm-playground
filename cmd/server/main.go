package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/prnvbn/textfsm-playground/internal/server"
)

func main() {
	// TODO? make this configurable
	cfg := &server.Config{
		Port:    9999,
		WasmDir: "./wasm/dist",
	}

	srv := server.New(cfg)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv.Run(ctx)
}
