package shop

import (
	"net/http"
	"strconv"

	"github.com/enghasib/laundry_service/utils"
)

func (h *ShopHandler) GetAllShopsHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for retrieving all shops
	query := r.URL.Query()
	limitParam := query.Get("limit")
	pageParam := query.Get("page")

	limit := 10 // default limit
	page := 1   // default page

	if limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil {
			limit = l
		}
	}

	if pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil {
			page = p
		}
	}

	shops, err := h.Srv.List(limit, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendResponseWithPagination(w, shops, page, limit, 0)
}
