package service

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/repository"
	"errors"
)

func GetAllExpenses(userID uint) (expenses []models.Expense, err error) {
	expenses, err = repository.GetAllExpenses(userID)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func GetExpenseByID(userID, expenseID uint) (expense models.Expense, err error) {
	expense, err = repository.GetExpenseByID(userID, expenseID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return expense, errs.ErrOperationNotFound
		}
		return models.Expense{}, err
	}
	return expense, nil
}

func CreateExpense(expense models.Expense) error {
	if err := repository.CreateExpense(expense); err != nil {
		return err
	}
	return nil
}

func UpdateExpense(expense models.Expense) error {
	if err := repository.UpdateExpense(expense); err != nil {
		return err
	}
	return nil
}

func DeleteExpense(expenseID uint, userID uint) error {
	if err := repository.DeleteExpense(expenseID, userID); err != nil {
		return err
	}
	return nil
}
