package middlewares

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"time"
)

func WithMiddlewares(r chi.Router) {

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))
}
