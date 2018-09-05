package services

import "demo/models"

func GetAdsByPositionId(id int) ([]*models.Ad, error) {
	var list []*models.Ad
	err := models.Orm.Where("ad_position_id = ?", id).Desc("id").Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetAds(page, pageSize int) ([]*models.Ad, error) {
	ads := make([]*models.Ad, 0, pageSize)
	return ads, models.Orm.Limit(pageSize, (page-1)*pageSize).Desc("id").Find(&ads)
}
