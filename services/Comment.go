package services

import "demo/models"

func SaveComment(comment *models.Comment) (int64, error) {
	return  models.Orm.Insert(&comment)
}

func GetCommentCountBy(userId,typeId,valueId int) (int64, error) {
	return models.Orm.Where("type_id = ? and value_id = ? and user_id = ?", typeId, valueId, userId).Count(&models.Comment{})
}

func GetCommentCount(typeId,valueId int) (*map[string]interface{}) {
	allCount, _ := models.Orm.Where("type_id = ? and value_id = ?", typeId, valueId).Count(&models.Comment{})

	return &map[string]interface{}{"allCount": allCount, "hasPicCount": 0}
}

func GetComments(valueId,userId,typeId int) ([]*models.Comment, error) {
	var list []*models.Comment
	err := models.Orm.Where("type_id = ? and value_id = ? and user_id = ?", typeId, valueId, userId).Find(&list)
	return list, err
}