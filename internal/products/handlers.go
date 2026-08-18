package products

import (
	"net/http"

	"github.com/Shriom-Trivedi/ecom-go/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. CALL the service -> ListProduct
	// 2. Return JSON in an HTTP response

	products := []string{"Hello", "Products"}

	json.Write(w, http.StatusOK, products)
}
