package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/enghasib/laundry_service/config"
	"github.com/enghasib/laundry_service/utils"

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

	// swagger ui integration
	mux.Handle("/ping", http.HandlerFunc(PingHandler))
	utils.RegisterSwaggerMux(mux)

	serverPort := ":" + strconv.Itoa(svr.cnf.HttpPort)
	fmt.Println("Server is running on port:5500")
	if err := http.ListenAndServe(serverPort, wrappedMuxWitMiddleware); err != nil {
		fmt.Println("Server starting failed!")
		os.Exit(1)
	}
}

// PingHandler
// @Summary Ping endpoint
// @Description This endpoint is used to check the health of the server. It responds with a simple "pong" message.
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}
