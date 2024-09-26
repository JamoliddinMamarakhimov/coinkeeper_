package db

import "coinkeeper/models"

func Migrate() error {
	err := dbConn.AutoMigrate(models.User{},
		models.Income{},
		models.Outcome{},
		models.Expense{},
		models.Card{},
	)
	if err != nil {
		return err
	}
	return nil
}
