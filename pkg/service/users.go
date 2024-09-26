package service

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/repository"
	"coinkeeper/utils"
	"errors"
)

func CreateUser(user models.User) error {
	userFromDB, err := repository.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}
	if userFromDB.ID > 0 {
		return errs.ErrUsernameUniquenessFailed
	}

	user.Password = utils.GenerateHash(user.Password)

	err = repository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return user, errs.ErrRecordNotFound
		}
		return user, err
	}
	return user, nil
}

func DeleteUser(id uint) error {
	err := repository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
