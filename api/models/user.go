package models

import (
	"api/utils/token"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null" json:"email" binding:"required"`
	Password string `gorm:"unique;not null" json:"password" binding:"required"`
}

func GetUserByID(id string) (*User, error) {
	var user User
	err := DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email, password string) (string, error) {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return err.Error(), err
	}
	err = VerifyPassword(user.Password, password)
	if err != nil {
		return err.Error(), err
	}
	token, err := token.GenerateToken(int(user.ID))
	if err != nil {
		return err.Error(), err
	}
	return token, err
}

func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {
	// turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}
