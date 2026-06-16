package entity

// time equivale ao Date em js/ts
// biblioteca do próprio go
import "time"

type Product struct {
	ID          string
	SKU         string
	Slug        string
	Name        string
	Brand       string
	Description string
	Price       float64
	Volume      int
	ImageURL    *string
	Stock       int
	FragranceID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
