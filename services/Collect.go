package services

import "demo/models"

func GetCollectByTypeId(userId,typeId int) (*models.Collect, error) {
	var model models.Collect
	has, err := models.Orm.Where("user_id = ? and type_id = ?", userId, typeId).Get(&model)
	if err != nil {
		return nil, err
	}
	if has {
		var good *models.Goods
		hasGood, _ := models.Orm.Id(model.ValueId).Get(&good)
		if hasGood {
			model.GoodsName = good.Name
			model.GoodsPicUrl = good.ListPicUrl
			model.GoodsBrief = good.GoodsBrief
			model.GoodsPrice = good.RetailPrice
		}
	}
	return &model, nil
}

func SaveCollect(collect *models.Collect) (string, error) {
	retStr := "add"
	has, err := models.Orm.Get(&collect)
	if err != nil {
		return "", err
	}
	if !has {
		models.Orm.Insert(&collect)
	} else {
		models.Orm.Delete(&collect)
		retStr = "delete"
	}
	return retStr, nil
}

func IsUserHasCollect(goodsId,userId,typeId int)  (int64, error) {
	return models.Orm.Where("value_id = ? and user_id = ? and type_id = ?", goodsId, userId,typeId).Count(&models.Collect{})
}