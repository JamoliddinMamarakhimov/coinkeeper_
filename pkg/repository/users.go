package repository

import (
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/models"
)

func CreateUser(user models.User) (err error) {
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Println("[repository.CreateUser] cannot create user. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Find(&users).Error
	if err != nil {
		logger.Error.Println("[repository.GetAllUsers] cannot get all users. Error is:", err.Error())
		return nil, translateError(err)
	}
	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Println("[repository.GetUserByID] cannot get user by id. Error is:", err.Error())
		return user, translateError(err)
	}
	return user, nil
}

func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Println("[repository.GetUserByID] cannot get user by id. Error is:", err.Error())
		return user, translateError(err)
	}
	return user, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Println("[repository.GetUserByID] cannot get user by id. Error is:", err.Error())
		return user, translateError(err)
	}
	return user, nil
}

func UpdateUser(user models.User) error {
	err := db.GetDBConn().Save(&user).Error
	if err != nil {
		logger.Error.Println("[repository.UpdateUser] cannot update user. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func DeleteUser(id uint) (err error) {
	err = db.GetDBConn().
		Table("users").
		Where("id = ?", id).
		Update("is_deleted", true).Error
	if err != nil {
		logger.Error.Println("[repository.DeleteUser] cannot delete user. Error is:", err.Error())
		return err
	}
	return nil
}
