package shop

import (
	"encoding/json"
	"net/http"

	"github.com/enghasib/laundry_service/domain"
	"github.com/enghasib/laundry_service/utils"
)

func (h *ShopHandler) CreateShopHandler(w http.ResponseWriter, r *http.Request) {
	var newShop domain.Shop

	err := json.NewDecoder(r.Body).Decode(&newShop)
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

	shop, err := h.Srv.Create(domain.Shop{
		Name:           newShop.Name,
		Location:       newShop.Location,
		ContactNumber:  newShop.ContactNumber,
		PaymentDetails: newShop.PaymentDetails,
		CreatedBy:      user.Uuid,
		ShopOwner:      newShop.ShopOwner,
	})

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// mount and encode with response
	utils.SendResponse(w, http.StatusCreated, shop)
}
