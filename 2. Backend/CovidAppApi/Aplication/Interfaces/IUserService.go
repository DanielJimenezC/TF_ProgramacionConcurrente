package services

import (
	entity "../../domain/entities"
	userResponse "../communication"
)

// IUserService Interface
type IUserService interface {
	Create(user entity.User) error
	GetAll() ([]userResponse.UserResponse, error)
	Validate(user *entity.User) error
	GetByID(id int) (userResponse.UserResponse, error)
	Delete(id int) error
	Update(id int, user entity.User) error
	Login(user entity.User) (bool, error)
}
