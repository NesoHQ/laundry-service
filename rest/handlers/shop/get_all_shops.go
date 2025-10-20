package shop

import (
	"net/http"
	"strconv"

	"github.com/enghasib/laundry_service/utils"
)

// GetAllShopsHandler godoc
// @Summary Get all shops
// @Description Retrieve a list of all shops with pagination
// @Tags Shop
// @Accept  json
// @Produce  json
// @Param limit query int false "Number of shops to return" default(10)
// @Param page query int false "Page number" default(1)
// @Success 200 {object} []domain.Shop
// @Param Authorization header string true "Bearer"
// @Failure 400 {object} map[string]string
// @Router /shops [get]
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
