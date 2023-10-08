package service

import (
	"server/config"
	"server/models"
)

func Get(u *[]models.User) (err error) {
	if err = config.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func Create(u *models.User) (err error) {
	if err = config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func Update(u *models.User, id string) (err error) {
	if err = config.DB.Model(&u).Where("id = ?", id).Update(u).Error; err != nil {
		return err
	}
	return nil
}

func Destroy(u *models.User, id string) (err error) {
	if err = config.DB.Model(&u).Delete(u, id).Error; err != nil {
		return err
	}
	return nil
}
