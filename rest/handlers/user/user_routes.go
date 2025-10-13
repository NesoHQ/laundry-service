package user

import (
	"net/http"

	middleware "github.com/enghasib/laundry_service/rest/middlewares"
)

func (h *UserHandler) UserRoute(mux *http.ServeMux, manager *middleware.MiddlewareManager) *http.ServeMux {
	mux.Handle("POST /users/register", http.HandlerFunc(h.CreateUserHandler))
	mux.Handle("POST /users/login", http.HandlerFunc(h.LoginUser))

	mux.Handle("GET /users", manager.With(http.HandlerFunc(h.GetAllUserHandler)))

	return mux
}
