package server

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *Server) init() {
	s.Use(middleware.Logger)
	s.Use(middleware.Recoverer)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	s.Use(corsMiddleware.Handler)

	s.Use(metricsMiddleware)
	s.Handle("/metrics", promhttp.Handler())

	api := humachi.New(s, huma.DefaultConfig("My API", "1.0.0"))

	huma.Get(api, "/api/{name}", s.NameHandler)
}
