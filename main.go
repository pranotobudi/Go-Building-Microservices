package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/pranotobudi/Go-Building-Microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)
	p := handlers.NewProducts(l)
	// sm := http.NewServeMux()
	sm := mux.NewRouter()
	getSubRouter := sm.Methods(http.MethodGet).Subrouter()
	getSubRouter.HandleFunc("/", p.GetProducts)

	putSubRouter := sm.Methods(http.MethodPut).Subrouter()
	putSubRouter.HandleFunc("/{id:[0-9]+}", p.UpdateProduct)
	putSubRouter.Use(p.MiddlewareProductValidation)

	postSubRouter := sm.Methods(http.MethodPost).Subrouter()
	postSubRouter.HandleFunc("/", p.AddProducts)
	postSubRouter.Use(p.MiddlewareProductValidation)
	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gh)
	// sm.Handle("/products", p)

	fmt.Println("bismillah")
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}

	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Printf("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	// http.ListenAndServe(":8080", sm)
	// http.ListenAndServe(":80", sm)
}
