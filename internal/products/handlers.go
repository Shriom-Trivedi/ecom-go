package products

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Shriom-Trivedi/ecom-go/internal/json"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
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

func (h *handler) ListProductByID(w http.ResponseWriter, r *http.Request) {
	// calling service
	idstr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idstr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	product, err := h.service.ListProductByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return JSON response
	json.Write(w, http.StatusOK, product)
}
