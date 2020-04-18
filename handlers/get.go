package handlers

import (
	"net/http"

	"github.com/ewhite1/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
// 	 200: productResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// shouldn't ever get a error here but log it just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /product/{id} proudcts listSingleProduct
// Return a list of products from the database
// reponses:
// 	 200: productResponse
// 	 404: errorRepsonse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)
	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[Error] fetching product", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}
