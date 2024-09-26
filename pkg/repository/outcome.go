package repository

import (
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/models"
	"errors"
	"gorm.io/gorm"
)

func GetAllOutcome(userID uint, query string) ([]models.Outcome, error) {
	var outcome []models.Outcome

	query = "%" + query + "%"

	err := db.GetDBConn().Model(&models.Outcome{}).
		Joins("JOIN users ON users.id = outcomes.user_id").
		Joins("JOIN outcome_categories ON outcome_categories.id = outcomes.category_id").
		Where("outcomes.user_id = ? AND outcomes.is_deleted = false AND outcomes.description iLIKE ?", userID, query).
		Order("outcomes.id").
		Find(&outcome).Error

	//err := db.GetDBConn().Model(&models.Outcome{}).
	//	Joins("JOIN users ON users.id = outcomes.user_id").
	//	Where("outcomes.user_id = ? AND description iLIKE ?", userID, query).
	//	Order("outcomes.id").
	//	Find(&outcome).Error
	if err != nil {
		logger.Error.Println("[repository.GetAllOutcome] cannot get all outcome. Error is:", err.Error())
		return nil, translateError(err)
	}
	return outcome, nil

	//err = db.GetDBConn().Find(&outcome).Error
	//if err != nil {
	//	logger.Error.Println("[repository.GetAllOutcome] cannot get all outcome. Error is:", err.Error())
	//	return nil, err
	//}
	//return outcome, nil
}

func GetOutcomeByID(userID, outcomeID uint) (models.Outcome, error) {
	var outcome models.Outcome

	err := db.GetDBConn().Model(&models.Outcome{}).
		Joins("JOIN outcome_categories ON outcome_categories.id = outcomes.category_id").
		Where("outcomes.id = ? AND outcomes.user_id = ? AND outcomes.is_deleted = false", outcomeID, userID).
		First(&outcome).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {

	}
	if err != nil {
		logger.Error.Println("[repository.GetOutcomeByID] cannot get outcome by id. Error is:", err.Error())
		return models.Outcome{}, translateError(err)
	}

	return outcome, nil
}

//func GetOutcomeByID(userID, outcomeID uint) (outcome models.Outcome, err error) {
//	err = db.GetDBConn().Model(&models.Outcome{}).
//		Joins("JOIN users ON users.id = outcomes.user_id").
//		Where("outcomes.user_id = ? AND outcomes.id = ?", userID, outcomeID).
//		First(&outcome).Error
//	if err != nil {
//		logger.Error.Println("[repository.GetOutcomeByID] cannot get outcome by id. Error is:", err.Error())
//		return models.Outcome{}, translateError(err)
//	}
//	return outcome, nil
//}

func CreateOutcome(outcome models.Outcome) error {
	err := db.GetDBConn().Create(&outcome).Error
	if err != nil {
		logger.Error.Println("[repository.CreateOutcome] cannot create outcome. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func UpdateOutcome(outcome models.Outcome) error {
	err := db.GetDBConn().Model(&outcome).Where("id = ?", outcome.ID).Save(outcome).Error

	//err := db.GetDBConn().Save(&outcome).Error
	if err != nil {
		logger.Error.Println("[repository.UpdateOutcome] cannot update outcome. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func DeleteOutcome(outcomeID, userID uint) error {
	// Обновляем флаг is_deleted на true

	err := db.GetDBConn().Exec("UPDATE outcomes set is_deleted = true WHERE id = $1 AND user_id = $2", outcomeID, userID).Error
	if err != nil {
		logger.Error.Println("[repository.DeleteOutcome] cannot delete outcome. Error is:", err.Error())
		return translateError(err)
	}

	return nil
}

//func DeleteOutcome(outcomeID int, userID uint) error {
//	err := db.GetDBConn().
//		Table("outcomes").
//		Where("id = ? AND user_id = ?", outcomeID, userID).
//		Update("is_deleted", true).Error
//	if err != nil {
//		logger.Error.Println("[repository.DeleteOutcome] cannot delete outcome. Error is:", err.Error())
//		return err
//	}
//	return nil
//
//}
