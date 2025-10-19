package user

import (
	"github.com/enghasib/laundry_service/config"
	middleware "github.com/enghasib/laundry_service/rest/middlewares"
)

type UserHandler struct {
	middleware middleware.Middlewares
	cnf        *config.Config
	srv        UserService
}

func NewUserHandler(
	middlewares *middleware.Middlewares,
	cnf *config.Config,
	srv UserService,
) *UserHandler {
	return &UserHandler{
		middleware: *middlewares,
		cnf:        cnf,
		srv:        srv,
	}
}
