package repository

import "github.com/dayli-fragrance/api/internal/domain/entity"

type FragranceRepository interface {
	FindAll() ([]entity.Fragrance, error)
	FindById(id string) (*entity.Fragrance, error)
}
