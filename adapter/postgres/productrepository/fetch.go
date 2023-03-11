package productrepository

import (
	"clean-architecture/core/domain"
	"clean-architecture/core/dto"
	"context"
	"github.com/booscaaa/go-paginate/paginate"
)

func (repository repository) Fetch(pagination *dto.PaginationParamRequest) (*domain.Pagination[[]domain.Product], error) {
	ctx := context.Background()
	var products []domain.Product
	total := int32(0)

	pagin := paginate.Instance(pagination)
	query, countQuery := pagin.
		Query("SELECT * FROM").
		Sort(pagination.Sort).
		Desc(pagination.Descending).
		Page(pagination.Page).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search).
		Select()

	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			product := domain.Product{}

			rows.Scan(
				&product.ID,
				&product.Name,
				&product.Price,
				&product.Description,
			)

			products = append(products, product)
		}
	}

	{
		err := repository.db.QueryRow(ctx, *countQuery).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	var p = domain.Pagination[[]domain.Product]

	p.Items = products
	p.Total = total
	return &p, nil
}
