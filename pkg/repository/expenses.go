package repository

import (
	"coinkeeper/db"
	"coinkeeper/models"
)

func GetAllExpenses(userID uint) ([]models.Expense, error) {
	var expenses []models.Expense
	err := db.GetDBConn().Where("user_id = ?", userID).Find(&expenses).Error
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func GetExpenseByID(userID, expenseID uint) (models.Expense, error) {
	var expense models.Expense
	err := db.GetDBConn().Where("id = ? AND user_id = ?", expenseID, userID).First(&expense).Error
	if err != nil {
		return models.Expense{}, err
	}
	return expense, nil
}

func CreateExpense(expense models.Expense) error {
	err := db.GetDBConn().Create(&expense).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateExpense(expense models.Expense) error {
	err := db.GetDBConn().Save(&expense).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteExpense(expenseID uint, userID uint) error {
	err := db.GetDBConn().Model(&models.Expense{}).
		Where("id = ? AND user_id = ?", expenseID, userID).
		Update("is_deleted", true).Error
	if err != nil {
		return err
	}
	return nil
}
