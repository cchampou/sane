package userHandler

import (
	"../lib/db"
	"../models"
	"log"
	"net/http"
)

type Handler struct {
	db db.DBLayer
}

func NewHandler() (*Handler, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	return &Handler{
		db: db,
	}, nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if h.db == nil {
		log.Fatal("Fail to bind db reference")
	}
	h.db.CreateUser(userModel.User{
		Email:     "clement@champouillon.com",
		Firstname: "Clément",
		Name:      "Champouillon",
	})
	w.Write([]byte("user created"))
}
