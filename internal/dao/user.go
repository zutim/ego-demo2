package dao

import (
	"ego-demo2/internal/entity"
	"github.com/zutim/ego/component"
)

type userDao struct {
	baseDao
}

func User() *userDao {
	return &userDao{}
}

func (dao *userDao) FindAll() ([]*entity.UserEntity, error) {
	var user []*entity.UserEntity
	db := component.Provider().Gorm()

	res := db.Table("user").Find(&user)
	return user, res.Error
}

// GetByUsername 根据用户名获取记录
func (dao *userDao) GetByEmail(email string) (*entity.UserEntity, error) {
	query := dao.db.Table(entity.TableUser).
		Where("email = ?", email).
		Where(entity.SoftDeleteCondition)

	user := new(entity.UserEntity)
	if err := query.First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
