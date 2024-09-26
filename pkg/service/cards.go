package service

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/repository"
	"errors"
	"gorm.io/gorm"
)

func GetAllCards(userID uint) (cards []models.Card, err error) {
	cards, err = repository.GetAllCards(userID)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func GetCardByID(userID, cardID uint) (card models.Card, err error) {
	card, err = repository.GetCardByID(userID, cardID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return card, errs.ErrOperationNotFound
		}
		return models.Card{}, err
	}
	return card, nil
}

func CreateCard(card models.Card) error {
	if err := repository.CreateCard(card); err != nil {
		return err
	}
	return nil
}

func UpdateCardBalance(cardID uint, amount float32) error {
	if err := repository.UpdateCardBalance(cardID, amount); err != nil {
		return err
	}
	return nil
}

func DeleteCard(cardID, userID uint) error {
	if err := repository.DeleteCard(cardID, userID); err != nil {
		return err
	}
	return nil
}
