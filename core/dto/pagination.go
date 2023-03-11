package dto

import (
	"net/http"
	"strconv"
	"strings"
)

type PaginationParamRequest struct {
	Search       string   `json:"search"`
	Descending   []string `json:"descending"`
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"ItemsPerPage"`
	Sort         []string `json:"sort"`
}

func FromValuePaginationRequestParams(request *http.Request) (*PaginationParamRequest, error) {
	page, _ := strconv.Atoi(request.FormValue("page"))
	itemsPerPage, _ := strconv.Atoi(request.FormValue("ItemsPerPage"))

	paginationRequestParam := PaginationParamRequest{
		Search:       request.FormValue("search"),
		Descending:   strings.Split(request.FormValue("descending"), ","),
		Sort:         strings.Split(request.FormValue("sort"), ","),
		Page:         page,
		ItemsPerPage: itemsPerPage,
	}

	return &paginationRequestParam, nil
}
