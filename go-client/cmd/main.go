package main

import (
	"log"
	"net/http"

	controller "github.com/gabscristofani/go-client/internal/controller/action"
)

func main() {
	http.HandleFunc("/cotacao", controller.Handle)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
