package shop

import "github.com/enghasib/laundry_service/domain"

type ShopService interface {
	Create(shop domain.Shop) (*domain.Shop, error)
	Get(shopId string) (*domain.Shop, error)
	List(limit, page int) ([]*domain.Shop, error)
	Update(shopId string, Shop domain.Shop) (*domain.Shop, error)
	Delete(shopId string) error
}
