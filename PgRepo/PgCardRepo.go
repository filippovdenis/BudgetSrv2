package PgRepo

import (
	"BudgetSrv2/models"

	"github.com/jinzhu/gorm"
)

type PgCardRepo struct {
	Db *gorm.DB
}

func (repo PgCardRepo) GetCards() []models.Card {
	var cards []models.Card
	repo.Db.Debug().Find(&cards)
	return cards
}

func (repo PgCardRepo) GetCardByName(CardKey string) *models.Card {

	println(CardKey)
	var card models.Card

	repo.Db.Debug().Where("card_name = ?", CardKey).Find(&card)

	return &card
}

func (repo PgCardRepo) InsertCard(card *models.Card) {
	println(card.CardName)
	println(card.Description)
	repo.Db.Debug().Create(card)
}

func (repo PgCardRepo) UpdateCardByName(CardName string, card *models.Card) {

	pgcard := &PgCard{}

	repo.Db.Debug().Where("card_name = ?", CardName).Debug().Find(pgcard)
	if pgcard.CardName == "" {
		println("Not Found")
		return
	}

	pgcard.Description = card.Description

	println(pgcard.CardName)
	repo.Db.Debug().Save(pgcard)
}

func (repo PgCardRepo) DeleteCardByName(CardName string) {
	println(CardName)
	repo.Db.Table("cards").Debug().Where("card_name= ?", CardName).Debug().Delete(&models.Card{})
}