package PgRepo

type PgCard struct {
	CardId int `gorm:"primary_key"`
	CardName string
	Description string
}

func (PgCard) TableName() string {
	return "cards"
}