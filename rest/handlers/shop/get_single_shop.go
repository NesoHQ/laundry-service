package shop

import (
	"net/http"

	"github.com/enghasib/laundry_service/utils"
)

// @Summary Get a single shop
// @Description Retrieve a shop by its ID
// @Tags Shop
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <token>"
// @Param shop_id path string true "Shop ID"
// @Success 200 {object} domain.Shop
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /shops/{shop_id} [get]
func (h *ShopHandler) GetSingleShopHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for retrieving a single shop by ID
	shopID := r.URL.Path[len("/shops/"):]
	if shopID == "" {
		http.Error(w, "Shop ID is required", http.StatusBadRequest)
		return
	}

	shop, err := h.Srv.Get(shopID)
	if err != nil {
		http.Error(w, "Shop not found", http.StatusNotFound)
		return
	}

	utils.SendResponse(w, http.StatusOK, shop)
}
