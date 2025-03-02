package usecase

import (
    "github.com/LucasPaz-7/Secretaria_Api_Go/model"
    "github.com/LucasPaz-7/Secretaria_Api_Go/repository"
    "errors"
)

type UserUseCase struct {
    repo *repository.UserRepository
}

func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
    return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Login(email, password string) (*model.User, error) {
    user, err := uc.repo.FindUserByEmailAndPassword(email, password)
    
    if err != nil {
        return nil, errors.New("credenciais inválidas")
    }
    
    return user, nil
}