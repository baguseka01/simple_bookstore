package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	FirtsName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Password  []byte `gorm:"column:-"`
}

func (user *User) HashPassword(password string) error {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashPassword

	return nil
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
