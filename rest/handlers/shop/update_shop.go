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
}

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
	})
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendResponse(w, http.StatusOK, shop)
}
