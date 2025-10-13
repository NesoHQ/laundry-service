package user

import "github.com/enghasib/laundry_service/domain"

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (srv *service) Create(user domain.User) (*domain.User, error) {
	usr, err := srv.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (srv *service) Get(userId int) (*domain.User, error) {
	user, err := srv.userRepo.Get(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (srv *service) List(limit, page int) ([]*domain.User, error) {
	listOfUser, err := srv.userRepo.List(limit, page)
	if err != nil {
		return nil, err
	}
	return listOfUser, nil
}

func (srv *service) Update(id int, user domain.User) (*domain.User, error) {
	usr, err := srv.userRepo.Update(id, user)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (srv *service) Delete(userId int) error {
	if err := srv.userRepo.Delete(userId); err != nil {
		return err
	}
	return nil
}

func (srv *service) Find(email, password string) (*domain.User, error) {
	user, err := srv.userRepo.Find(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
