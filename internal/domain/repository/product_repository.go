package repository

import "github.com/dayli-fragrance/api/internal/domain/entity"

// * = ponteiro, é um endereço de memória pode ou não ser nil (null)

type ProductRepository interface {
	FindAll() ([]entity.Product, error)
	FindBySlug(slug string) (*entity.Product, error)
}