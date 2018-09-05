package services

import "demo/models"

func GetUserById(id int64) (*models.User, error) {
	var user models.User
	has, err := models.Orm.Id(id).Get(&user)
	//has, err := Orm.Get(&user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, models.ErrNotExist
	}
	return &user, nil
}