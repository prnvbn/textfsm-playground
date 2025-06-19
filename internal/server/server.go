package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Server struct {
	*chi.Mux
	cfg *Config
}

func New(cfg *Config) *Server {
	router := chi.NewRouter()
	s := &Server{
		Mux: router,
		cfg: cfg,
	}

	s.init()
	return s
}

func (s *Server) Run(ctx context.Context) {
	addr := fmt.Sprintf("0.0.0.0:%d", s.cfg.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: s,
	}

	go func() {
		log.Info().Str("addr", addr).Msg("serving API")
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
