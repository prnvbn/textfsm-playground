package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Server struct {
	*chi.Mux
	cfg     *Config
	wasmDir string
}

func New(cfg *Config) *Server {
	router := chi.NewRouter()
	s := &Server{
		Mux:     router,
		cfg:     cfg,
		wasmDir: cfg.WasmDir,
	}

	s.init()
	return s
}

func (s *Server) init() {
	fs := http.FileServer(http.Dir(s.wasmDir))
	s.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(w, r)
	})
}

func (s *Server) Run(ctx context.Context) {
	addr := fmt.Sprintf("0.0.0.0:%d", s.cfg.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: s,
	}

	go func() {
		log.Info().Str("addr", addr).Str("wasm_dir", s.wasmDir).Msg("serving API and WASM files")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("server error")
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info().Msg("attempting graceful shutdown")

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("graceful shutdown failed")
		if err := srv.Close(); err != nil {
			log.Error().Err(err).Msg("failed to stop server")
		}
	}
}
