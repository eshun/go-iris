package services

import (
	"demo/models"
	"strconv"
)

func GetCategorys() ([]*models.Category, error) {
	var list []*models.Category
	err := models.Orm.Where("parent_id = ? and name <> ?", 0, "推荐").Cols("id", "name").Find(&list)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		for i, v := range list {
			//所属栏目
			ids :=GetChildCategoryIdsByParentId(v.Id)
			goodsList, _ := GetGoodsByCategoryIds(ids)
			list[i].GoodsList = goodsList
		}
	}
	return list, nil
}

func GetCategoryLimit(limit int) ([]*models.Category, error) {
	var list []*models.Category
	err := models.Orm.Limit(limit, 0).Find(&list)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		for i, v := range list {
			models.Orm.Where("parent_id = ?", v.Id).Find(&list[i].SubCategoryList)
		}
	}
	return list, nil
}

func GetCategoryById(id int) (*models.Category, error) {
	var model models.Category
	_, err := models.Orm.Where("id = ?", id).Get(&model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func GetCategoryByIds(ids string) ([]*models.Category, error) {
	var data []*models.Category
	err := models.Orm.Where("id in ("+ids+")").Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetCategoryAndChildById(id int) (*models.Category, error) {
	var model models.Category
	has, err := models.Orm.Where("id = ?", id).Get(&model)
	if err != nil {
		return nil, err
	}
	if has {
		//model.SubCategoryList=
		models.Orm.Where("parent_id = ?", model.Id).Find(&model.SubCategoryList)
	}
	return &model, nil
}

func GetChildCategoryByParentId(parentId int) ([]*models.Category, error) {
	var list []*models.Category
	err := models.Orm.Where("parent_id = ?", parentId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetChildCategoryIdsByParentId(parentId int) string {
	//.Where("id in (" + article.Relation + ")")
	var ids string
	var list []*models.Category
	models.Orm.Where("parent_id = ?", parentId).Cols("id").Find(&list)
	if len(list) > 0 {
		for i, v := range list {
			if i > 0 {
				ids += ","
			}
			ids += strconv.Itoa(v.Id)
		}
	}
	return ids
}

func GetParentCategoryIdsByIds(ids string) string {
	var retIds string
	var list []*models.Category
	models.Orm.Where("id in ("+ids+")").Cols("parent_id").Find(&list)
	if len(list) > 0 {
		for i, v := range list {
			if i > 0 {
				retIds += ","
			}
			retIds += strconv.Itoa(v.ParentId)
		}
	}
	return retIds
}