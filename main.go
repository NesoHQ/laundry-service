package main

import (
	"github.com/enghasib/laundry_service/cmd"
)

//	@title			Laundry Service Backend API
//	@version		1.0
//	@description	This is a backend API for a laundry service application. It provides endpoints for managing users, orders, and laundry services. The API is built using Go and follows RESTful principles. It includes features such as user authentication, order processing, and service management. The API is designed to be scalable and maintainable, making it suitable for integration with various frontend applications.

//	@contact.name	API Support
//	@contact.url	https://github.com/NesoHQ/laundry-service
//	@termsOfService	https://github.com/NesoHQ/laundry-service

//	@license.name	Apache 2.0

//	@host		localhost:5500
//	@BasePath	/

func main() {
	cmd.Serve()
}
