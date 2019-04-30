package models

type Shop struct {
	ShopName string `gorm:"primary_key"`
	ArticleName string
}
