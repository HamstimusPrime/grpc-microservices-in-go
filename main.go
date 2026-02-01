package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HamstimusPrime/grpc-microservices-in-go/handlers"
)

func main() {
	port := "9090"
	l := log.New(os.Stdout, "product-api-", log.LstdFlags)
	ph := handlers.NewProductsHandler(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      sm,
		IdleTimeout:  2 * time.Minute,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	go func() {
		fmt.Printf("server running on port: %v\n", port)
		log.Fatal(server.ListenAndServe())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("quitting server")

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
