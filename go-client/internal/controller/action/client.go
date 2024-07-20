package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gabscristofani/go-client/internal/controller/response"
	"github.com/gabscristofani/go-client/internal/model"
)

var (
	timeout = 5 * time.Second
)

func Handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	log.Println("Request received")

	var c model.Cambio

	if err := c.NewCambioServer(ctx); err != nil {
		log.Printf("Error getting exchange rate: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := c.SaveCotacao(); err != nil {
		log.Printf("Error saving exchange rate: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.NewSuccess(c, http.StatusOK).Send(w)
}
