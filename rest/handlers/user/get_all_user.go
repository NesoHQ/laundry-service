package user

import (
	"net/http"
	"strconv"

	"github.com/enghasib/laundry_service/utils"
)

// GetAllUserHandler handles the retrieval of all users with pagination.
// @Summary Get all users
// @Description Retrieve a paginated list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} []domain.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *UserHandler) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	page, _ := strconv.Atoi(queryParam.Get("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(queryParam.Get("limit"))
	if limit == 0 {
		limit = 10
	}
	userList, err := h.srv.List(limit, page)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponseWithPagination(w, userList, page, limit, 0)
}
