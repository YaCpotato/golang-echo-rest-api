package repository

import "goapi/app/infrastructure/mysql/entity"

type IUserRepository interface {
	FindByID(id uint) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id uint) error
}