package usersRouter

import (
	"../../handlers"
	"github.com/go-chi/chi"
	"log"
)

func Routes(r chi.Router) {
	h, err := userHandler.NewHandler()
	if err != nil {
		log.Fatal(err)
	}
	r.Get("/", h.Create)
}
