package services

import (
	"errors"

	entity "../entities"
	model "../entities"
	userRepository "../repositories"
)

type serv struct{}

var userRepo userRepository.IUserRepository = userRepository.UserRepository()

// UserService Implemantation
func UserService() IUserService {
	return &serv{}
}

func (*serv) Validate(user *model.User) error {
	if user == nil {
		err := errors.New("The Request Body is empty")
		return err
	}
	if user.Username == "" {
		err := errors.New("The username is empty")
		return err
	}
	if user.Password == "" {
		err := errors.New("The password is empty")
		return err
	}
	return nil
}

func (*serv) Create(user model.User) error {
	err := userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (*serv) GetAll() ([]model.User, error) {
	response, err := userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (*serv) GetByID(id int) (entity.User, error) {
	response, err := userRepo.GetUser(id)
	if err != nil {
		return response, err
	}
	return response, nil
}
