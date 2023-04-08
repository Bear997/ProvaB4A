package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Location struct {
	Latitude  string `json:"latitude" binding:"required" validate:"required"`
	Longitude string `json:"longitude" binding:"required" validate:"required"`
}

type Card struct {
	gorm.Model
	Title    string   `json:"title" binding:"required" validate:"required"`
	Image    string   `json:"image"`
	Position Location `json:"position" binding:"required" validate:"required" gorm:"embedded"`
	Users    []User   `gorm:"many2many:user_cards;"`
}

func CreateCard(db *gorm.DB, Card *Card) (err error) {

	err = db.Create(Card).Error
	if err != nil {
		fmt.Println("sto in model errore")
		return err
	}
	return nil
}

func GetCardFromPosition(db *gorm.DB, Card *Card, lat string, lon string) (err error) {

	const radius = 0.2
	err = db.Where("6371 * acos(cos(radians(?)) * cos(radians(latitude + 0)) * cos(radians(longitude + 0) - radians(?)) + sin(radians(?)) * sin(radians(latitude + 0))) <= ?", lat, lon, lat, radius).First(Card).Error

	// err = db.Where("latitude = ?", lat).Where("longitude = ?", lon).First(Card).Error

	if err != nil {
		println(err)
		return err
	}
	return nil
}

func GetCardFromId(db *gorm.DB, card *Card, id string) (err error) {
	err = db.First(card, id).Error
	if err != nil {
		return err
	}
	return nil
}

func AssignCardToUser(db *gorm.DB, association *UserCard) (err error) {
	fmt.Println("qui ci arrivo sto in model assign")

	err = db.Create(association).Error

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
