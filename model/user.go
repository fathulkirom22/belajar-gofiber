package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required" gorm:"unique"`
	Password string `json:"password" validate:"required"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	if user.Password != "" {
		hash, err := hashPassword(user.Password)
		if err != nil {
			return err
		}
		db.Model(user).Where(&User{Email: user.Email}).Update("Password", hash)
		return nil
	}
	return errors.New("password null")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
