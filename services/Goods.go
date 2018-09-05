package services

import (
	"demo/models"
	"strconv"
)

func GetGoodById(id int) models.Goods {
	model := models.Goods{}
	models.Orm.Id(id).Get(&model)
	return model
}

func GetNewGoods() ([]*models.Goods, error) {
	var list []*models.Goods
	err := models.Orm.Where("is_new = ?", 1).Cols("id", "name", "list_pic_url", "retail_price").Limit(4, 0).Desc("id").Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetHotGoods() ([]*models.Goods, error) {
	var list []*models.Goods
	err := models.Orm.Where("is_hot = ?", 1).Cols("id", "name", "list_pic_url", "retail_price", "goods_brief").Limit(3, 0).Desc("id").Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetGoodsByCategoryIds(ids string) ([]*models.Goods, error) {
	var list []*models.Goods
	err := models.Orm.Where("category_id in (" + ids + ")").Cols("id", "name", "list_pic_url", "retail_price").Limit(7, 0).Desc("id").Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetAllGoods() ([]*models.Goods, error) {
	var data []*models.Goods
	err := models.Orm.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetGoodsBy(strOrder,strWhere string,limit,start int) ([]*models.Goods, error) {
	var data []*models.Goods
	err := models.Orm.Cols("id, name, list_pic_url, retail_price").Where(strWhere).OrderBy(strOrder).Limit(limit, start).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetProductsByGoodsId(goodsId int) ([]*models.Product, error) {
	var data []*models.Product
	err := models.Orm.Where("goods_id = ?", goodsId).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetSpecificationListByGoodsId(goodsId int) ([]*models.GoodsSpecification, error) {
	var data []*models.GoodsSpecification
	err := models.Orm.Where("goods_id = ?", goodsId).Find(&data)
	if err != nil {
		return nil, err
	}
	if len(data) > 0 {
		for i, v := range data {
			var spec *models.Specification
			has, _ := models.Orm.Id(v.SpecificationId).Get(&spec)
			if has {
				data[i].Name = ""
			}
		}
	}
	return data, nil
}

func GetSaleCount() (int64, error) {
	return models.Orm.Where("is_delete = ? and is_on_sale = ?", 0, 1).Count(&models.Goods{})
}

func GetRelatedGoodsIdsByGoodsId(goodsId int) string {
	var ids string
	var list []*models.RelatedGoods
	models.Orm.Where("goods_id = ?", goodsId).Cols("related_goods_id").Find(&list)
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

func GetSameCategoryGoods(id int) ([]*models.Goods, error) {
	var good *models.Goods
	has, err := models.Orm.Id(id).Get(&good)
	if has {
		var goods []*models.Goods
		models.Orm.Where("category_id = ?", good.CategoryId).Find(&goods)
		return goods, nil

	} else {
		return nil, err
	}
}

func GetGoodsByIds(ids string) ([]*models.Goods, error) {
	var goods []*models.Goods
	err := models.Orm.Where("id in( ? )", ids).Find(&goods)
	return goods, err
}

func GetRelatedGoodsById(id int) ([]*models.Goods, error) {
	ids := GetRelatedGoodsIdsByGoodsId(id)
	if ids != "" {
		return GetGoodsByIds(ids)
	} else {
		return GetSameCategoryGoods(id)
	}
}

func GetCategoryIdsBy(strWhere string) string {
	var ids string
	var list []*models.Goods
	models.Orm.Where(strWhere).Cols("category_id").Find(&list)
	if len(list) > 0 {
		for i, v := range list {
			if i > 0 {
				ids += ","
			}
			ids += strconv.Itoa(v.CategoryId)
		}
	}
	return ids
}

func GetGoodsGalleryByGoodsId(goodsId int) ([]*models.GoodsGallery, error) {
	var list []*models.GoodsGallery
	err := models.Orm.Where("goods_id = ?", goodsId).Find(&list)
	return list, err
}

func GetGoodsAttributeByGoodsId(goodsId int) ([]*models.GoodsAttribute, error) {
	var list []*models.GoodsAttribute
	err := models.Orm.Where("goods_id = ?", goodsId).Find(&list)
	if len(list) > 0 {
		for i, v := range list {
			var att *models.Attribute
			has, _ := models.Orm.Id(v.AttributeId).Get(&att)
			if has {
				list[i].Name = att.Name
			}
		}
	}
	return list, err
}

func GetBrandBId(id int) (*models.Brand, error) {
	var model *models.Brand
	_,err:=models.Orm.Id(id).Get(&model)
	return model,err
}

func GetIssueByGoodsId(goodsId int) ([]*models.GoodsIssue, error) {
	var list []*models.GoodsIssue
	err := models.Orm.Where("goods_id = ?", goodsId).Find(&list)
	return list, err
}
