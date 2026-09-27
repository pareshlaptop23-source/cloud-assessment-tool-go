package repository

import (
	"cloud-assessment-tool/config"
	"cloud-assessment-tool/models"
)

func CreateUser(user *models.User) error {

	return config.DB.Create(user).Error

}

func GetUserByEmail(email string) (models.User, error) {

	var user models.User

	err := config.DB.
		Where("email=?", email).
		First(&user).Error

	return user, err

}

func GetAllUsers() ([]models.User, error) {

	var users []models.User

	err := config.DB.Find(&users).Error

	return users, err

}
