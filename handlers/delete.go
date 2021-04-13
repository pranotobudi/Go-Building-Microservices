package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pranotobudi/Go-Building-Microservices/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// response:
// 201: noContent

// DeleteProduct returns the products from the data store
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("handle DELETE product", id)
	err := data.DeleteProduct(id)
	if err != data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
