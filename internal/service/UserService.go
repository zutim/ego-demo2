package service

import (
	"ego-demo2/internal/entity"
	"ego-demo2/internal/repository"
	"ego-demo2/internal/repository/impl"
)

type userService struct {
	userRep repository.UserRepo
}

var user *userService

func GetUser() *userService {
	once.Do(func() {
		user = &userService{
			userRep: impl.NewUserRepo(),
		}
	})
	return user
}

func (u *userService) GetAll() ([]*entity.UserEntity, error) {
	return u.userRep.FindAll()
}
