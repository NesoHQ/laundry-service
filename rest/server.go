package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/enghasib/laundry_service/config"
	"github.com/enghasib/laundry_service/rest/handlers/shop"
	"github.com/enghasib/laundry_service/rest/handlers/user"
	middleware "github.com/enghasib/laundry_service/rest/middlewares"
)

type server struct {
	cnf         *config.Config
	userHandler user.UserHandler
	shopHandler shop.ShopHandler
}

func NewServer(cnf config.Config, userHandler user.UserHandler, shopHandler shop.ShopHandler) *server {
	return &server{
		cnf:         &cnf,
		userHandler: userHandler,
		shopHandler: shopHandler,
	}
}

func (svr *server) Start() {
	mux := http.NewServeMux()

	middlewareManager := middleware.NewMiddlewareManager()
	middlewareManager.Use(middleware.Logger, middleware.Cors)

	wrappedMuxWitMiddleware := middlewareManager.Apply(mux)

	svr.userHandler.UserRoute(mux, middlewareManager)
	svr.shopHandler.ShopRoute(mux, middlewareManager)

	serverPort := ":" + strconv.Itoa(svr.cnf.HttpPort)
	fmt.Println("Server is running on port:5500")
	if err := http.ListenAndServe(serverPort, wrappedMuxWitMiddleware); err != nil {
		fmt.Println("Server starting failed!")
		os.Exit(1)
	}
}
