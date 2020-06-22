package services

import (
	"errors"

	model "../../domain/entities"
	interfaces "../../domain/interfaces"
	repository "../../infraestructure/repositories"
	service "../interfaces"
)

type serv struct{}

var userRepo interfaces.IUserRepository = repository.UserRepository()

// UserService Implemantation
func UserService() service.IUserService {
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

func (*serv) GetByID(id int) (model.User, error) {
	response, err := userRepo.GetUser(id)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (*serv) Delete(id int) error {
	err := userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (*serv) Update(id int, user model.User) error {
	err := userRepo.Update(id, user)
	if err != nil {
		return err
	}
	return nil
}
