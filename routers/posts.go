package postsRouter

import (
	"github.com/go-chi/chi"
	"net/http"
)

func Routes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
}
