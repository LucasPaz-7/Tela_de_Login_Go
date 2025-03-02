package repository

import (
	"github.com/LucasPaz-7/Secretaria_Api_Go/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindUserByEmailAndPassword(email, password string) (*model.User, error) {
    var user model.User
    result := r.DB.Where("email = ? AND password = ?", email, password).First(&user)
    return &user, result.Error
}