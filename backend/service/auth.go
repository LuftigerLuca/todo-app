package service

import (
	"errors"
	"todo-app/dto"
	"todo-app/model"
	"todo-app/util"
)

type AuthService struct{}

func (s *AuthService) Register(req dto.RegisterRequest) error {
	u := model.User{
		Username: req.Username,
		Password: req.Password,
	}
	return model.DB.Create(&u).Error
}

func (s *AuthService) Login(req dto.LoginRequest) (string, error) {
	var u model.User
	err := model.DB.Where("username = ?", req.Username).Take(&u).Error
	if err != nil {
		return "", errors.New("username or password is incorrect")
	}

	if err := model.VerifyPassword(req.Password, u.Password); err != nil {
		return "", errors.New("username or password is incorrect")
	}

	return util.GenerateToken(u.ID)
}
