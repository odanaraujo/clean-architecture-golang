package di

import (
	"clean-architecture/adapter/http/productservice"
	"clean-architecture/adapter/postgres"
	"clean-architecture/adapter/postgres/productrepository"
	"clean-architecture/core/domain"
	"clean-architecture/core/domain/usecase/productusecase"
)

func ConficProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)
	return productService
}
