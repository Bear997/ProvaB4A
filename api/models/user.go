package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required" validate:"required" gorm:"unique"`
	Password string `json:"password" binding:"required" validate:"required"`
	Role     string `json:"role"  validate:"required"`
}

func CreateUser(db *gorm.DB, user *User) (err error) {
	err = db.Create(user).Error
	if err != nil {
		fmt.Println("sto in model errore")
		return err
	}
	return nil
}

func GetUserFromId(db *gorm.DB, user *User, id string) (err error) {
	err = db.First(user, id).Error
	if err != nil {
		return err
	}
	return nil
}

func Login(db *gorm.DB, user *User, email string) (err error) {

	err = db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
