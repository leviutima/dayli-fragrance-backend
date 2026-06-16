package usecase

import (
	"github.com/dayli-fragrance/api/internal/domain/entity"
	"github.com/dayli-fragrance/api/internal/domain/repository"
)

// Instancia do repo, constructor to ts: constructor(private repo: Repository){}
type GetProductsUseCase struct {
	repo repository.ProductRepository
}

func NewGetProductsUseCase(repo repository.ProductRepository) *GetProductsUseCase {
	return &GetProductsUseCase{repo: repo}
}


								//parametro de envio	// parametro de resposta
								
func (useCase *GetProductsUseCase) Execute()           ([]entity.Product, error) {
	return useCase.repo.FindAll()
}
