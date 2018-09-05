package services

import "demo/models"

func GetBrands() ([]*models.Brand, error) {
	var list []*models.Brand
	err := models.Orm.Where("is_new = ?", 1).Limit(4, 0).Asc("new_sort_order").Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetBrandById(id int) (*models.Brand, error) {
	var model models.Brand
	has, err := models.Orm.Id(id).Get(&model)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, models.ErrNotExist
	}
	return &model, nil
}

func GetPageBrands(limit,start int) ([]*models.Brand, error) {
	var list  []*models.Brand
	err := models.Orm.Cols("id", "name", "floor_price", "app_list_pic_url").Limit(limit, start).Desc("id").Find(&list)
	return list, err
}