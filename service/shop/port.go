package shop

import (
	"github.com/enghasib/laundry_service/domain"
	shopHandler "github.com/enghasib/laundry_service/rest/handlers/shop"
)

type ShopService interface {
	shopHandler.ShopService
}

type ShopRepo interface {
	Create(shop domain.Shop) (*domain.Shop, error)
	Get(shopId string) (*domain.Shop, error)
	List(limit, page int) ([]*domain.Shop, error)
	Update(shopId string, Shop domain.Shop) (*domain.Shop, error)
	Delete(shopId string) error
}
