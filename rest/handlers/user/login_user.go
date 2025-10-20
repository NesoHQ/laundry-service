package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/enghasib/laundry_service/utils"
)

type LoginCredential struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type loginResponse struct {
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

// @Summary User login
// @Description Authenticate a user and return a JWT token.
// @Tags Users
// @Accept json
// @Produce json
// @Param request body LoginCredential true "User login credentials"
// @Success 200 {object} loginResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/login [post]
func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	var requestBody LoginCredential
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		fmt.Println("error:", err.Error())
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	user, err := h.srv.Find(requestBody.Email, requestBody.Password)
	if err != nil {
		http.Error(w, "invalid credential", http.StatusNotFound)
		return
	}

	payload := utils.Payload{
		Uuid:     user.Uuid,
		UserName: user.UserName,
		Email:    user.Email,
		Role:     user.Role,
	}

	jwt, err := utils.CreateToken(h.cnf.JwtSecretKey, payload)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := loginResponse{
		Message:     "User Login successfully!",
		AccessToken: jwt,
	}

	utils.SendResponse(w, http.StatusOK, response)
}
