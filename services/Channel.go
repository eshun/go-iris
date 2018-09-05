package services

import "demo/models"

func GetChannels() ([]*models.Channel, error) {
	var list []*models.Channel
	err := models.Orm.Asc("sort_order").Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}