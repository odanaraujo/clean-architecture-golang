package productrepository

import (
	"clean-architecture/core/domain"
	"clean-architecture/core/dto"
	"context"
)

func (repository repository) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	ctx := context.Background()
	product := domain.Product{}

	err := repository.db.QueryRow(ctx,
		"INSERT INTO product (name, price, description) VALUES ($1, $2, $3) returning *",
		productRequest.Name,
		productRequest.Price,
		productRequest.Description,
	).Scan(
		&product.Name,
		&product.Price,
		&product.Description)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
