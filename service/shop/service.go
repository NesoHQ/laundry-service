package shop

import "github.com/enghasib/laundry_service/domain"

type shopService struct {
	shopRepo ShopRepo
}

func NewShopService(shopRepo ShopRepo) ShopService {
	return &shopService{
		shopRepo: shopRepo,
	}
}

func (srv *shopService) Create(shop domain.Shop) (*domain.Shop, error) {
	creteShop, err := srv.shopRepo.Create(shop)
	if err != nil {
		return nil, err
	}
	return creteShop, err
}

func (srv *shopService) Get(shopId string) (*domain.Shop, error) {
	shop, err := srv.shopRepo.Get(shopId)
	if err != nil {
		return nil, err
	}
	return shop, err
}
func (srv *shopService) List(limit, page int) ([]*domain.Shop, error) {
	shops, err := srv.shopRepo.List(limit, page)
	if err != nil {
		return nil, err
	}
	return shops, err
}
func (srv *shopService) Delete(shopId string) error {
	err := srv.shopRepo.Delete(shopId)
	if err != nil {
		return err
	}
	return nil
}
func (srv *shopService) Update(shopId string, Shop domain.Shop) (*domain.Shop, error) {
	updatedShop, err := srv.shopRepo.Update(shopId, Shop)
	if err != nil {
		return nil, err
	}
	return updatedShop, nil
}
