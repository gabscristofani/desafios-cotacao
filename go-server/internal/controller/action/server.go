package controller

import (
	"log"
	"net/http"

	"github.com/gabscristofani/go-server/internal/controller/response"
	"github.com/gabscristofani/go-server/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var c model.Cambio
	if err := c.NewTaxaCambio(); err != nil {
		log.Printf("Error getting exchange rate: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := c.SaveCotacao(); err != nil {
		log.Printf("Error saving exchange rate: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// response.NewSuccess(c, http.StatusOK
	response.NewSuccess(c, http.StatusOK).Send(w)
}
