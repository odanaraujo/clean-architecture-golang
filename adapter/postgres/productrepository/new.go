package productrepository

import (
	"clean-architecture/adapter/postgres"
	"clean-architecture/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{db: db}
}
