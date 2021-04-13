// Package classification of product API
//
// Documentation for Product API
//
// Scheme: http
// BasePath: /
// Version: 1.0.0
// Consumes:
// - Application/json
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pranotobudi/Go-Building-Microservices/data"
)

// A list of products return in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type productNocontent struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// in: path
	// required:true
	ID int `json:"id"`
}
type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.AddProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT ", r.URL.Path)
		regex := regexp.MustCompile(`/([0-9]+)`)
		path := r.URL.Path
		g := regex.FindAllStringSubmatch(path, -1)
		p.l.Println("g = ", g)
		if len(g) != 1 {
			p.l.Println("Invalid len(g)")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 { // --> domain.com/1, the result should be [/1 1], len == 2
			p.l.Println("Invalid len(g[0])")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := g[0][1] // --> domain.com/1, the result should be [/1 1], we'll take 1
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Unable to convert to number")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		p.l.Printf("got ID: %d", id)
		// p.UpdateProduct(id, rw, r)
		return
	}

	http.Error(rw, "ERROR METHOD", http.StatusMethodNotAllowed)
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product data: ")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	// prod := &data.Product{}
	// err := prod.DecodeJSON(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)
	// }
	p.l.Printf("Prod: %#v \n", prod)
	data.AddProduct(prod)
}

// func (p *Products) UpdateProduct(id int, rw http.ResponseWriter, r *http.Request) {
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Update Product ")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("Handle PUT Update Product 1 ")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Println("Handle PUT Update Product 2 ")
	err := data.UpdateProduct(id, prod)
	p.l.Println("Handle PUT Update Product 3")
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	p.l.Println("Handle MIDDLEWARE Product ")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		p.l.Println("Handle MIDDLEWARE Product 1")
		prod := &data.Product{}
		p.l.Println("Handle MIDDLEWARE Product 2")
		err := prod.DecodeJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			return
		}
		err = prod.Validate()
		if err != nil {
			p.l.Println("Unable to validate product", err)
			http.Error(rw, fmt.Sprintf("Unable to validate product: %s", err), http.StatusBadRequest)
			return
		}
		p.l.Println("Handle MIDDLEWARE Product 3")
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		r.Context()
		next.ServeHTTP(rw, req)
		p.l.Println("Handle MIDDLEWARE Product 4")

	})
}
