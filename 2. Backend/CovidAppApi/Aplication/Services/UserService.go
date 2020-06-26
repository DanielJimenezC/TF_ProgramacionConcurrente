package services

import (
	"errors"

	model "../../domain/entities"
	interfaces "../../domain/interfaces"
	repository "../../infraestructure/repositories"
	userResponse "../communication"
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

func (*serv) GetAll() ([]userResponse.UserResponse, error) {
	response, err := userRepo.GetAll()
	users := make([]userResponse.UserResponse, 0)

	for i := 0; i < len(response); i++ {
		var user userResponse.UserResponse
		user.ID = response[i].ID
		user.Username = response[i].Username
		users = append(users, user)
	}

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (*serv) GetByID(id int) (userResponse.UserResponse, error) {
	response, err := userRepo.GetUser(id)
	var user userResponse.UserResponse
	user.ID = response.ID
	user.Username = response.Username
	if err != nil {
		return user, err
	}
	return user, nil
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

func (*serv) Login(user model.User) (bool, error) {
	login, _ := userRepo.Login(user)
	if !login {
		return false, nil
	}
	return true, nil
}
