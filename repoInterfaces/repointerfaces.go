package repoInterfaces

import "BudgetSrv2/models"

type CardRepo interface {

	GetCards() []models.Card

	GetCardByName(string) *models.Card

	InsertCard(*models.Card)

	UpdateCardByName(string, *models.Card)

	DeleteCardByName(string)
}