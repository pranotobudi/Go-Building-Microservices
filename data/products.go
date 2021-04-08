package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku,omitempty"`
	CreatedOn   string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) DecodeJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

type Products []*Product

func (p *Products) EncodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}
func GetProducts() Products {
	return productList
}
func AddProduct(p *Product) {
	p.ID = GetNextID()
	productList = append(productList, p)
}
func UpdateProduct(id int, p *Product) error {
	idx, err := findProduct(id)
	if err != nil {
		return err
	}
	productList[idx] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product Not Found")

func findProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ErrProductNotFound
}
func GetNextID() int {
	prod := productList[len(productList)-1]
	return prod.ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
