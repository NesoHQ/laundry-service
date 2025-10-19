package user

import "github.com/enghasib/laundry_service/domain"

type userService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (srv *userService) Create(user domain.User) (*domain.User, error) {
	usr, err := srv.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (srv *userService) Get(userId string) (*domain.User, error) {
	user, err := srv.userRepo.Get(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (srv *userService) List(limit, page int) ([]*domain.User, error) {
	listOfUser, err := srv.userRepo.List(limit, page)
	if err != nil {
		return nil, err
	}
	return listOfUser, nil
}

func (srv *userService) Update(userId string, user domain.User) (*domain.User, error) {
	usr, err := srv.userRepo.Update(userId, user)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (srv *userService) Delete(userId string) error {
	if err := srv.userRepo.Delete(userId); err != nil {
		return err
	}
	return nil
}

func (srv *userService) Find(email, password string) (*domain.User, error) {
	user, err := srv.userRepo.Find(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
