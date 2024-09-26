package repository

import (
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/models"
)

func CreateCard(card models.Card) error {
	err := db.GetDBConn().Create(&card).Error
	if err != nil {
		logger.Error.Println("[repository.CreateCard] cannot create card. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func UpdateCardBalance(cardID uint, amount float32) error {
	var card models.Card

	// Находим карту
	if err := db.GetDBConn().First(&card, cardID).Error; err != nil {
		logger.Error.Println("[repository.UpdateCardBalance] cannot find card. Error is:", err.Error())
		return translateError(err)
	}

	// Обновляем баланс
	card.Balance += amount
	if err := db.GetDBConn().Save(&card).Error; err != nil {
		logger.Error.Println("[repository.UpdateCardBalance] cannot update card balance. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func GetAllCards(userID uint) ([]models.Card, error) {
	var cards []models.Card
	if err := db.GetDBConn().Where("user_id = ?", userID).Find(&cards).Error; err != nil {
		logger.Error.Println("[repository.GetAllCards] cannot find card. Error is:", err.Error())
		return nil, err
	}
	return cards, nil
}

func GetCardByID(userID, cardID uint) (models.Card, error) {
	var card models.Card
	err := db.GetDBConn().Where("id = ? AND user_id = ?", cardID, userID).First(&card).Error
	if err != nil {
		return models.Card{}, err
	}
	return card, nil
}

func DeleteCard(cardID, userID uint) error {
	err := db.GetDBConn().Model(&models.Card{}).
		Where("id = ?", cardID).
		Update("is_deleted", true).Error
	if err != nil {
		logger.Error.Println("[service.DeleteCard] cannot delete card. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}
