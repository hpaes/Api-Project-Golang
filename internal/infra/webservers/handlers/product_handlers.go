package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hpaes/api-project-golang/internal/dto"
	"github.com/hpaes/api-project-golang/internal/entity"
	"github.com/hpaes/api-project-golang/internal/infra/database"
)

type ProductHandler struct {
	productDb database.ProductInterface
}

func NewProductHandler(productDb database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		productDb: productDb,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// USECASE
	p, error := entity.NewProduct(product.Name, product.Description, product.Price)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	if err := h.productDb.Create(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
