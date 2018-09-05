package services

import (
	"demo/models"
)

func SaveFootprint(goodsId,userId int) (*models.Footprint, error) {
	model := models.Footprint{UserId: userId, GoodsId: goodsId}
	_, err := models.Orm.Insert(model)
	return &model, err
}

func DeleFootprintById(userId int) (int64, error) {
	model := new(models.Footprint)
	return models.Orm.Where("user_id = ?", userId).Delete(&model)
}