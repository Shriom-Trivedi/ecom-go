package orders

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

// TODO: Do some payload validation here
func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {

	var tempOrder createOrderParams
	if err := json.Read(r, &tempOrder); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdOrder, err := h.service.PlaceOrder(r.Context(), tempOrder)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusCreated, createdOrder)
}
