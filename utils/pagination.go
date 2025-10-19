package utils

import (
	"encoding/json"
	"net/http"
)

type PaginatedData struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}

func SendResponseWithPagination(w http.ResponseWriter, data any, page int, limit int, count int) {

	totalPage := 0

	if count > limit {
		totalPage = count / limit
	}

	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page:      page,
			Limit:     limit,
			TotalPage: totalPage,
			TotalData: count,
		},
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(paginatedData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
