package service

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/repository"
	"errors"
)

func GetAllOutcome(userID uint, query string) (outcome []models.Outcome, err error) {
	outcome, err = repository.GetAllOutcome(userID, query)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func GetOutcomeByID(userID, outcomeID uint) (outcome models.Outcome, err error) {
	outcome, err = repository.GetOutcomeByID(userID, outcomeID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return outcome, errs.ErrOperationNotFound
		}
		return outcome, err
	}
	return outcome, nil
}

func CreateOutcome(outcome models.Outcome) error {
	if err := repository.CreateOutcome(outcome); err != nil {
		return err
	}
	return nil
}

func UpdateOutcome(outcome models.Outcome) error {
	if err := repository.UpdateOutcome(outcome); err != nil {
		return err
	}
	return nil
}

func DeleteOutcome(outcomeID int, userID uint) error {
	if err := repository.DeleteOutcome(uint(outcomeID), userID); err != nil {
		return err
	}
	return nil
}
