package shop

import (
	"encoding/json"
	"net/http"

	"github.com/enghasib/laundry_service/domain"
	"github.com/enghasib/laundry_service/utils"
)

type createShopRequest struct {
	Name           string `json:"name" validate:"required"`
	Location       string `json:"location" validate:"required"`
	ContactNumber  string `json:"contact_number" validate:"required"`
	PaymentDetails string `json:"payment_details" validate:"required"`
	ShopOwner      string `json:"shop_owner"`
	CoverImage     string `json:"cover_image"`
}

// CreateShopHandler godoc
// @Summary Create a new shop
// @Description Create a shop with necessary details
// @Tags Shop
// @Accept  json
// @Produce  json
// @Success 201 {object} createShopRequest
// @Param Authorization header string true "Bearer <token>"
// @Param request body createShopRequest true "Shop to create"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /shops [post]
func (h *ShopHandler) CreateShopHandler(w http.ResponseWriter, r *http.Request) {
	var newShop createShopRequest
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
		CoverImage:     newShop.CoverImage,
	})
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}
	// mount and encode with response
	utils.SendResponse(w, http.StatusCreated, shop)
}
