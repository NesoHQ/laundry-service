package cmd

import (
	"fmt"
	"os"

	"github.com/enghasib/laundry_service/config"
	appDB "github.com/enghasib/laundry_service/infra/db"
	userRepo "github.com/enghasib/laundry_service/repo/user"
	"github.com/enghasib/laundry_service/rest"
	user_handler "github.com/enghasib/laundry_service/rest/handlers/user"
	middleware "github.com/enghasib/laundry_service/rest/middlewares"
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

	//service
	userService := user_service.NewService(usrRepo)

	// handler
	userHandler := user_handler.NewUserHandler(middlewares, cnf, userService)

	server := rest.NewServer(*cnf, *userHandler)
	server.Start()

}
