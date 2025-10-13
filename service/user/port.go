package user

import (
	"github.com/enghasib/laundry_service/domain"
	userHandler "github.com/enghasib/laundry_service/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(User domain.User) (*domain.User, error)
	Get(userId int) (*domain.User, error)
	List(limit, page int) ([]*domain.User, error)
	Update(id int, User domain.User) (*domain.User, error)
	Delete(userId int) error
	Find(email, password string) (*domain.User, error)
}
