package userModel

import (
	"net/http"
)

type user struct {
	id        int
	email     string
	firstname string
	name      string
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user created"))
}
