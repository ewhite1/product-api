package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product define the structure for the API, with a Staticly defined data before we introduce a db
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// toJSON serializes the conents of the collection to JSON. Provides better performance than json.Unmarshal.
// This reduces the allocation and the overhead of the service

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// writes the products to the db
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// returns a list of products
func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)

}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

//generate a id for products
func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// hard coded list of products for now before a db is used
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "a Frothy milk coffee",
		Price:       3.50,
		SKU:         "abcd221",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Epresso",
		Description: "A small, strong coffee without milk",
		Price:       2.00,
		SKU:         "fddd221",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
