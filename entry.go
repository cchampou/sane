package main

import (
	"./lib"
	"./routers"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	middlewares.WithMiddlewares(r)

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Route("/posts", func(r chi.Router) {
		postsRouter.Routes(r)
	})

	http.ListenAndServe(":3000", r)
}
