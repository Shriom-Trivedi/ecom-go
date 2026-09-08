package products

import (
	"log"
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
	// 1. CALL the service -> ListProducts

	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Return JSON in an HTTP response

	json.Write(w, http.StatusOK, products)
}
