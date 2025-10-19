package user

import (
	"github.com/enghasib/laundry_service/domain"
	userHandler "github.com/enghasib/laundry_service/rest/handlers/user"
)

type UserService interface {
	userHandler.UserService
}

type UserRepo interface {
	Create(User domain.User) (*domain.User, error)
	Get(userId string) (*domain.User, error)
	List(limit, page int) ([]*domain.User, error)
	Update(userId string, User domain.User) (*domain.User, error)
	Delete(userId string) error
	Find(email, password string) (*domain.User, error)
}
