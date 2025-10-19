package shop

import (
	"net/http"

	middleware "github.com/enghasib/laundry_service/rest/middlewares"
)

func (h *ShopHandler) ShopRoute(mux *http.ServeMux, manager *middleware.MiddlewareManager) *http.ServeMux {
	mux.Handle("POST /shops", manager.With(http.HandlerFunc(h.CreateShopHandler), h.Middleware.Authentication))
	mux.Handle("GET /shops", manager.With(http.HandlerFunc(h.GetAllShopsHandler), h.Middleware.Authentication))
	mux.Handle("GET /shops/{shop_id}", manager.With(http.HandlerFunc(h.GetSingleShopHandler)))
	mux.Handle("PUT /shops/{shop_id}", manager.With(http.HandlerFunc(h.UpdateShopHandler), h.Middleware.Authentication))
	mux.Handle("DELETE /shops/{shop_id}", manager.With(http.HandlerFunc(h.DeleteShopHandler), h.Middleware.Authentication))
	return mux
}
