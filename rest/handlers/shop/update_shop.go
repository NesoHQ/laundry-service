package shop

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/enghasib/laundry_service/domain"
	"github.com/enghasib/laundry_service/utils"
)

type shopUpdatableProps struct {
	Name           string `json:"name" db:"name" validate:"required"`
	Location       string `json:"location" db:"location"`
	ContactNumber  string `json:"contact" db:"contact"`
	PaymentDetails string `json:"payment_details" db:"payment_details"`
	ShopOwner      string `json:"shop_owner,omitempty" db:"shop_owner"`
}

// @Summary Update a shop
// @Description Update shop details by its ID
// @Tags Shop
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <token>"
// @Param shop_id path string true "Shop ID"
// @Param request body shopUpdatableProps true "Shop properties to update"
// @Success 200 {object} domain.Shop
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /shops/{shop_id} [put]
func (h *ShopHandler) UpdateShopHandler(w http.ResponseWriter, r *http.Request) {
	var updatedProps shopUpdatableProps

	shopId := r.URL.Path[len("/shops/"):]
	fmt.Println("shopid", shopId)

	if shopId == "" {
		utils.SendError(w, http.StatusBadRequest, "shop_id is required")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&updatedProps)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, ok := utils.GetUserFromContext(r, *h.Cnf)

	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if user.Role != "admin" {
		utils.SendError(w, http.StatusForbidden, "Forbidden: Admins only")
		return
	}

	shop, err := h.Srv.Update(shopId, domain.Shop{
		Name:           updatedProps.Name,
		Location:       updatedProps.Location,
		ContactNumber:  updatedProps.ContactNumber,
		PaymentDetails: updatedProps.PaymentDetails,
		ShopOwner:      updatedProps.ShopOwner,
	})
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendResponse(w, http.StatusOK, shop)
}
