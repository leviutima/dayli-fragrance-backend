package handler

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/dayli-fragrance/api/internal/application/usecase"
)

type ProductHandler struct {
	getProductsUseCase *usecase.GetProductsUseCase
}

func NewProductHandler(getProductsUseCase *usecase.GetProductsUseCase) *ProductHandler {
	return &ProductHandler{getProductsUseCase: getProductsUseCase}
}

func (handler *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var products, err = handler.getProductsUseCase.Execute()
	if err != nil {
		fmt.Println("Erro:", err)
		http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

