package repository

import (
	"github.com/dinhuy9595/shop/domain/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// UserRepository is
type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}

// NewUserRepository is
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

//
func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
