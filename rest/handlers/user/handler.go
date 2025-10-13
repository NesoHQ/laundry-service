package user

import (
	"github.com/enghasib/laundry_service/config"
	middleware "github.com/enghasib/laundry_service/rest/middlewares"
)

type UserHandler struct {
	middleware middleware.Middlewares
	cnf        *config.Config
	srv        Service
}

func NewUserHandler(
	middlewares *middleware.Middlewares,
	cnf *config.Config,
	srv Service,
) *UserHandler {
	return &UserHandler{
		middleware: *middlewares,
		cnf:        cnf,
		srv:        srv,
	}
}
