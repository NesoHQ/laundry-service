package shop

import (
	"net/http"

	"github.com/enghasib/laundry_service/utils"
)

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
