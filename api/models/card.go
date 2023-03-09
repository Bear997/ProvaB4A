package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Location = struct {
	latitude  string `gorm:"notnull"`
	longitude string `gorm:"notnull"`
}

type Card = struct {
	ID       int      `gorm:"primaryKey"`
	title    string   `gorm:"notnull"`
	location Location `gorm:"notnull"`
	verified bool     `gorm:"notnull"`
}

func CreateCard(db *gorm.DB, Card *Card) (err error) {
	err = db.Create(Card).Error
	if err != nil {
		fmt.Println("sto in model errore")
		return err
	}
	return nil
}
