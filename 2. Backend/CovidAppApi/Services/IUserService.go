package services

import entity "../entities"

// IUserService Interface
type IUserService interface {
	Create(user entity.User) error
	GetAll() ([]entity.User, error)
	Validate(user *entity.User) error
	GetByID(id int) (entity.User, error)
	Delete(id int) error
	Update(id int, user entity.User) error
}
