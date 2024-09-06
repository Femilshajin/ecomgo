package product

import (
	"fmt"
	"net/http"

	"github.com/femilshajin/ecomgo/types"
	"github.com/femilshajin/ecomgo/utils"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /", h.handleGetProducts)
	router.HandleFunc("POST /", h.handleCreateProduct)
	return router
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, ps)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.ProductPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	p := types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	}

	id, err := h.store.AddProduct(p)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	p.ID = id

	utils.WriteJson(w, http.StatusOK, p)
}
