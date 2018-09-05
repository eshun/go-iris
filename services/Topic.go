package services

import "demo/models"

func GetTopics(limit int) ([]*models.Topic, error) {
	var list []*models.Topic
	err := models.Orm.Limit(limit, 0).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetTopicsByCategoryId(categoryId,limit int) ([]*models.Topic, error) {
	var list []*models.Topic
	err := models.Orm.Where("is_show = ? and topic_category_id= ? ",1,categoryId).Limit(limit, 0).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetTopicById(id int) (*models.Topic, error) {
	model := models.Topic{}
	has, err := models.Orm.Where("id = ?", id).Get(&model)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, models.ErrNotExist
	}
	return &model, nil
}

func GetTopicCount() (int64, error){
	return models.Orm.Count(&models.Topic{})
}

func GetPageTopics(limit,start int) ([]*models.Topic, error) {
	var list  []*models.Topic
	err := models.Orm.Cols("id", "title", "subtitle", "price_info", "scene_pic_url").Limit(limit, start).Desc("id").Find(&list)
	return list, err
}