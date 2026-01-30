package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HamstimusPrime/grpc-microservices-in-go/handlers"
)

func main() {
	port := "9090"
	l := log.New(os.Stdout, "product-api-", log.LstdFlags)
	hh := handlers.NewHandler(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: sm,
	}
	fmt.Printf("server running on port: %v\n", port)
	log.Fatal(server.ListenAndServe())

}
