package handlers

import (
	"net/http"

	"github.com/pranotobudi/Go-Building-Microservices/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// response:
// 200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.EncodeJSON(rw)
	// jsonByte, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusInternalServerError)
	}
	// rw.Write(jsonByte)
	// enc := json.NewEncoder(rw)
	// enc.Encode(lp)
}
