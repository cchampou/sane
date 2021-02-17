package main

import (
	"./lib/middlewares"
	"./routers/posts"
	"./routers/users"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
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

	log.Fatal(http.ListenAndServe(":3000", r))
}
