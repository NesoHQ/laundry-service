package user

import (
	"encoding/json"
	"net/http"

	"github.com/enghasib/laundry_service/domain"
	"github.com/enghasib/laundry_service/utils"
)

type UserResponse struct {
	Message   string `json:"message"`
	ID        int    `json:"id"`
	Unique_id string `json:"unique_id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type createUserRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// create product
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var newUser createUserRequest

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.srv.Create(domain.User{
		UserName: newUser.UserName,
		Email:    newUser.Email,
		Password: newUser.Password,
	})

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// mount and encode with response
	utils.SendResponse(w, http.StatusCreated, UserResponse{
		Message:   "User created Successfully!",
		ID:        user.Id,
		Unique_id: user.Uuid,
		UserName:  user.UserName,
		Email:     user.Email,
		Role:      user.Role,
	})

}
