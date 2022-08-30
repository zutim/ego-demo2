package repository

import (
	"ego-demo2/internal/entity"
)

type UserRepo interface {
	FindAll() ([]*entity.UserEntity, error)
}
