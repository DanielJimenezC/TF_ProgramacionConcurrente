package repositories

import entity "../entities"

// IUserRepository Interface
type IUserRepository interface {
	Create(user entity.User) error
	GetAll() ([]entity.User, error)
	GetUser(id int) (entity.User, error)
	Delete(id int) error
	Update(id int, user entity.User) error
}
