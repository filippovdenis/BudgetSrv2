package models

type Card struct {
	CardName string `gorm:"primary_key"`
	Description string
}
