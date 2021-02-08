package main

import (
	"./lib/db"
	"./lib/middlewares"
	"./routers/posts"
	"./routers/users"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	db.Connect()
	r := chi.NewRouter()

	middlewares.WithMiddlewares(r)

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Route("/users", func(r chi.Router) {
		usersRouter.Routes(r)
	})

	r.Route("/posts", func(r chi.Router) {
		postsRouter.Routes(r)
	})

	http.ListenAndServe(":3000", r)
}
