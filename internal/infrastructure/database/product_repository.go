package database

import (
	"context"

	"github.com/dayli-fragrance/api/internal/domain/entity"
	"github.com/jackc/pgx/v5"
)

type ProductRepositoryPostgres struct {
	connection *pgx.Conn
}

func NewProductRepositoryPostgres(connection *pgx.Conn) *ProductRepositoryPostgres {
	return &ProductRepositoryPostgres{connection: connection}
}

func (repository *ProductRepositoryPostgres) FindAll() ([]entity.Product, error) {
	var rows, err = repository.connection.Query(
		context.Background(),
		"SELECT id, sku, slug, name, brand, description, price, volume, image_url, stock, fragrance_id, created_at, updated_at FROM products",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var product entity.Product
		var scanErr = rows.Scan(
			&product.ID, &product.SKU, &product.Slug, &product.Name, &product.Brand,
			&product.Description, &product.Price, &product.Volume, &product.ImageURL,
			&product.Stock, &product.FragranceID, &product.CreatedAt, &product.UpdatedAt,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		products = append(products, product)
	}

	return products, nil
}

func (repository *ProductRepositoryPostgres) FindBySlug(slug string) (*entity.Product, error) {
	var product entity.Product
	var err = repository.connection.QueryRow(
		context.Background(),
		"SELECT id, sku, slug, name, brand, description, price, volume, image_url, stock, fragrance_id, created_at, updated_at FROM products WHERE slug = $1",
		slug,
	).Scan(
		&product.ID, &product.SKU, &product.Slug, &product.Name, &product.Brand,
		&product.Description, &product.Price, &product.Volume, &product.ImageURL,
		&product.Stock, &product.FragranceID, &product.CreatedAt, &product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
