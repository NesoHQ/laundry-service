package utils

import (
	"net/http"

	_ "github.com/enghasib/laundry_service/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func SwaggerHandler() http.Handler {
	return httpSwagger.WrapHandler
}

func RegisterSwaggerMux(mux *http.ServeMux) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}
