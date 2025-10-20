package shop

import (
	"net/http"

	"github.com/enghasib/laundry_service/utils"
)

// @Summary Delete a shop
// @Description Delete a shop by its ID
// @Tags Shop
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <token>"
// @Param shop_id path string true "Shop ID"
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /shops/{shop_id} [delete]
func (h *ShopHandler) DeleteShopHandler(w http.ResponseWriter, r *http.Request) {
	shopID := r.URL.Path[len("/shops/"):]
	if shopID == "" {
		http.Error(w, "Shop ID is required", http.StatusBadRequest)
		return
	}

	// check user role
	user, ok := utils.GetUserFromContext(r, *h.Cnf)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if user.Role != "admin" {
		http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
		return
	}

	err := h.Srv.Delete(shopID)
	if err != nil {
		http.Error(w, "Failed to delete shop", http.StatusInternalServerError)
		return
	}

	deleteResponse := struct {
		Message string `json:"message"`
	}{
		Message: "Shop deleted successfully",
	}

	utils.SendResponse(w, http.StatusNoContent, deleteResponse)
}
