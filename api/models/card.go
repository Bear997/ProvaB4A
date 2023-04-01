package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Location struct {
	Latitude  string `json:"latitude" binding:"required"`
	Longitude string `json:"longitude" binding:"required"`
}

type Card struct {
	gorm.Model
	ID       int      `gorm:"primary_key;auto_increment;not_null"`
	Title    string   `json:"title" binding:"required"`
	Image    string   `json:"image" binding:"required"`
	Position Location `json:"position" binding:"required" gorm:"embedded"` //? trovare il modo per fare notnull, devo mettere embedd senno non funziona
	Verified bool     `json:"verified" binding:"required"`
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

	err = db.Where("latitude = ?", lat).Where("longitude = ?", lon).First(Card).Error

	if err != nil {
		println(err)
		return err
	}
	return nil
}

func GetCardFromId(db *gorm.DB, card *Card, id string) (err error) {
	fmt.Println(id)
	err = db.First(card, id).Error
	if err != nil {
		return err
	}
	return nil
}
