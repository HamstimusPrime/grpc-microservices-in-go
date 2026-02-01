package handlers

import (
	"log"
	"net/http"

	data "github.com/HamstimusPrime/grpc-microservices-in-go/models"
)

type Products struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to parse JSON", http.StatusInternalServerError)
	}
}
