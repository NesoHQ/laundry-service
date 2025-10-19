package cmd

import (
	"fmt"
	"os"

	"github.com/enghasib/laundry_service/config"
	appDB "github.com/enghasib/laundry_service/infra/db"
	"github.com/enghasib/laundry_service/repo/shop"
	userRepo "github.com/enghasib/laundry_service/repo/user"
	"github.com/enghasib/laundry_service/rest"
	shop_handler "github.com/enghasib/laundry_service/rest/handlers/shop"
	user_handler "github.com/enghasib/laundry_service/rest/handlers/user"
	middleware "github.com/enghasib/laundry_service/rest/middlewares"
	shop_service "github.com/enghasib/laundry_service/service/shop"
	user_service "github.com/enghasib/laundry_service/service/user"
)

func Serve() {
	cnf := config.GetConfig()

	// db configuration
	db, err := appDB.NewConnection(cnf)
	if err != nil {
		fmt.Println("DB connection Error, Error:", err)
		os.Exit(1)
	}
	appDB.Migrate(cnf)

	// middleware
	middlewares := middleware.NewMiddlewares(cnf)

	//repo
	usrRepo := userRepo.NewUserRepo(db)
	shopRepo := shop.NewShopRepo(db)

	//service
	userService := user_service.NewUserService(usrRepo)
	shopService := shop_service.NewShopService(shopRepo)

	// handler
	userHandler := user_handler.NewUserHandler(middlewares, cnf, userService)
	shopHandler := shop_handler.NewShopHandler(middlewares, cnf, shopService)

	server := rest.NewServer(*cnf, *userHandler, *shopHandler)
	server.Start()

}
