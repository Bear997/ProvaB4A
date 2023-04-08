package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required" validate:"email" binding:"email" gorm:"unique"`
	Password string `json:"password" binding:"required" validate:"required"`
	Nome     string `json:"nome" binding:"required" validate:"required"`
	Cognome  string `json:"cognome" binding:"required" validate:"required"`
	Role     string `json:"role" gorm:"default:user"`
	Cards    []Card `json:"cards" gorm:"many2many:userCards;"`
}

type UserCard struct {
	UserID   int  `gorm:"primaryKey"`
	CardID   int  `gorm:"primaryKey"`
	Verified bool `json:"verified"`
}

func CreateUser(db *gorm.DB, user *User) (err error) {
	err = db.Create(user).Error
	if err != nil {
		fmt.Println("sto in model errore")
		return err
	}
	return nil
}
func ChangeProfile(db *gorm.DB, user *User) (err error) {
	err = db.Model(&user).Updates(user).Error
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

func GetAllCardsOfUser(db *gorm.DB, id string, user *User, cards *[]Card) (err error) {

	err = db.Model(&user).Association("Cards").Find(&cards)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
