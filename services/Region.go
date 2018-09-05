package services

import "demo/models"

func GetRegionById(id int) (*models.Region, error) {
	var model models.Region
	has, err := models.Orm.Id(id).Get(&model)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, models.ErrNotExist
	}
	return &model, nil
}

func GetRegionByParentId(parentId int) ([]*models.Region, error) {
	var list []*models.Region
	err := models.Orm.Where("parent_id  = ?", parentId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}