package user

import "github.com/enghasib/laundry_service/domain"

type Service interface {
	Find(Email, Password string) (*domain.User, error)
	Create(user domain.User) (*domain.User, error)
	List(limit, page int) ([]*domain.User, error)
}
