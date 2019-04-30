package repoInterfaces

import "BudgetSrv2/models"

type CardRepo interface {

	getCards() []models.Card

	getCardByKey(string) *models.Card

	insertCard(*models.Card)

	updateCardByKey(string, *models.Card)

	deleteCardByKey(string)
}