package model

import (
	"html"
	"strings"
	"todo-app/util"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null; unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func VerifyPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {
	var err error

	u := User{}

	err = DB.Model(User{}).Where("username=?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil {
		return "", err
	}

	token, err := util.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}
