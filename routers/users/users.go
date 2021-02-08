package usersRouter

import (
	"../../models"
	"github.com/go-chi/chi"
)

func Routes(r chi.Router) {
	r.Get("/", userModel.Create)
}
