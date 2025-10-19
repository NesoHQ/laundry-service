package shop

import (
	"github.com/enghasib/laundry_service/config"
	middleware "github.com/enghasib/laundry_service/rest/middlewares"
)

type ShopHandler struct {
	Middleware middleware.Middlewares
	Cnf        *config.Config
	Srv        ShopService
}

func NewShopHandler(
	middlewares *middleware.Middlewares,
	Cnf *config.Config,
	Srv ShopService,
) *ShopHandler {
	return &ShopHandler{
		Middleware: *middlewares,
		Cnf:        Cnf,
		Srv:        Srv,
	}
}
