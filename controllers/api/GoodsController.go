package api

import (
	"demo/controllers"
	"demo/services"
	"strconv"
	"fmt"
)

type GoodsController struct{
	controllers.ApiController
}

func (c *GoodsController) Get() {
	data, err := services.GetAllGoods()
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(data)
	}
}

func (c *GoodsController) GetSkuBy(id int) {
	products, _ := services.GetProductsByGoodsId(id)
	specifications, _ := services.GetSpecificationListByGoodsId(id)

	c.Success(&map[string]interface{}{"specificationList": specifications, "productList": products})
}

func (c *GoodsController) GetCount() {
	count, err := services.GetSaleCount()
	if err != nil {
		c.Fail(err)
	} else{
		c.Success(&map[string]interface{}{"goodsCount": count})
	}
}

func (c *GoodsController) GetNew() {
	bannerInfo := map[string]interface{}{"url": "", "name": "坚持初心，为你寻觅世间好物", "img_url": "http://yanxuan.nosdn.127.net/8976116db321744084774643a933c5ce.png"}
	data := map[string]interface{}{"bannerInfo": bannerInfo}
	c.Success(&data)
}

func (c *GoodsController) GetHot() {
	bannerInfo := map[string]interface{}{"url": "", "name": "大家都在买的严选好物", "img_url": "http://yanxuan.nosdn.127.net/8976116db321744084774643a933c5ce.png"}
	data := map[string]interface{}{"bannerInfo": bannerInfo}
	c.Success(&data)
}

func (c *GoodsController) GetCategory() {
	id, err := c.Ctx.URLParamInt("id")
	currentCategory, err := services.GetCategoryById(id)
	if currentCategory != nil {
		parentCategory, _ := services.GetCategoryById(currentCategory.ParentId)
		brotherCategory, _ := services.GetChildCategoryByParentId(currentCategory.ParentId)
		c.Success(&map[string]interface{}{"currentCategory": currentCategory, "parentCategory": parentCategory, "brotherCategory": brotherCategory})
	} else {
		c.Fail(err)
	}
}

func (c *GoodsController) GetFilter() {
	strWhere := "1=1"
	categoryId, err := c.Ctx.URLParamInt("categoryId")
	if err == nil && categoryId>0 {
		categoryIds := services.GetChildCategoryIdsByParentId(categoryId)
		if categoryIds != "" {
			categoryIds = strconv.Itoa(categoryId) + "," + categoryIds
		} else {
			categoryIds = strconv.Itoa(categoryId)
		}
		strWhere += " and category_id in(" + categoryIds + ")"
	}
	hasNew :=c.Ctx.URLParamExists("isNew")
	if hasNew {
		isNew, _ := c.Ctx.URLParamInt("isNew")
		strWhere += " and is_new =" + strconv.Itoa(isNew)
	}
	hasHot :=c.Ctx.URLParamExists("isHot")
	if hasHot {
		isHot, _ := c.Ctx.URLParamInt("isHot")
		strWhere += " and is_hot =" + strconv.Itoa(isHot)
	}
	keyword := c.Ctx.URLParam("keyword")
	if keyword != "" {
		strWhere += " and name like '%" + keyword + "%'"
	}

	var filterCategory []interface{}

	ids := services.GetCategoryIdsBy(strWhere)
	if ids != "" {
		pids := services.GetParentCategoryIdsByIds(ids)
		data, _ := services.GetCategoryByIds(pids)
		if len(data) > 0 {
			m := make([]interface{}, len(data)+1)
			m[0] = map[string]interface{}{"id": 0, "name": "全部"}

			for i, v := range data {
				m[i+1] = map[string]interface{}{"id": v.Id, "name": v.Name}
			}
			filterCategory = m
		}
	}
	c.Success(filterCategory)
}

func (c *GoodsController) GetList() {
	strWhere := "1=1"
	strOrder := ""
	categoryId, err := c.Ctx.URLParamInt("categoryId")
	if err == nil && categoryId > 0 {
		categoryIds := services.GetChildCategoryIdsByParentId(categoryId)
		if categoryIds != "" {
			categoryIds = strconv.Itoa(categoryId) + "," + categoryIds
		} else {
			categoryIds = strconv.Itoa(categoryId)
		}
		strWhere += " and category_id in(" + categoryIds + ")"
	}
	brandId, err := c.Ctx.URLParamInt("brandId")
	if err == nil && brandId > 0 {
		strWhere += " and brand_id =" + strconv.Itoa(brandId)
	}
	hasNew := c.Ctx.URLParamExists("isNew")
	if hasNew {
		isNew, _ := c.Ctx.URLParamInt("isNew")
		strWhere += " and is_new =" + strconv.Itoa(isNew)
	}
	hasHot := c.Ctx.URLParamExists("isHot")
	if hasHot {
		isHot, _ := c.Ctx.URLParamInt("isHot")
		strWhere += " and is_hot =" + strconv.Itoa(isHot)
	}
	keyword := c.Ctx.URLParam("keyword")
	if keyword != "" {
		strWhere += " and name like '%" + keyword + "%'"
	}
	sort := c.Ctx.URLParam("sort")
	if sort == "price" {
		strOrder += " retail_price"
	} else {
		strOrder += " id"
	}
	order := c.Ctx.URLParam("order")
	if order == "desc" {
		strOrder += " desc"
	} else {
		strOrder += " asc"
	}

	ids := services.GetCategoryIdsBy(strWhere)
	//var filterCategory []interface{}
	var filterCategory []interface{}

	if ids != "" {
		pids := services.GetParentCategoryIdsByIds(ids)
		data, _ := services.GetCategoryByIds(pids)
		if len(data) > 0 {
			m := make([]interface{}, len(data)+1)
			m[0] = map[string]interface{}{"id": 0, "name": "全部", "checked": keyword == ""}

			for i, v := range data {
				m[i+1] = map[string]interface{}{"id": v.Id, "name": v.Name, "checked": v.Id == categoryId}
			}
			filterCategory = m
		}
	}
	fmt.Println(filterCategory)

	var start int
	size, err := c.Ctx.URLParamInt("size")
	if err != nil || size < 1 {
		size = 10
	}
	page, err := c.Ctx.URLParamInt("page")
	if err != nil || page < 2 {
		page = 1
		start = 0
	} else {
		start = size*(page-1) + 1
	}
	data, _ := services.GetGoodsBy(strOrder, strWhere, size, start)

	TotalPages := 1;
	Pagesize := size;
	CurrentPage := page;
	Count := len(data)
	count, err := services.GetSaleCount()
	if int(count) > size {
		TotalPages = (int(count) + 1) / size
	}

	data2 := map[string]interface{}{"count": Count, "totalPages": TotalPages, "pagesize": Pagesize, "currentPage": CurrentPage, "data": data, "filterCategory": filterCategory, "goodsList": data}
	c.Success(&data2)
}

func (c *GoodsController) GetDetail() {
	id, _ := c.Ctx.URLParamInt("id")
	good := services.GetGoodById(id)
	gallery, _ := services.GetGoodsGalleryByGoodsId(id)
	attribute, _ := services.GetGoodsAttributeByGoodsId(id)
	issue, _ := services.GetIssueByGoodsId(id)
	brand, _ := services.GetBrandBId(good.BrandId)

	commentCount, _ := services.GetCommentCountBy(id, 1, 0)

	//hotComment, err :=services.GetComments(id,1,0)

	comment := map[string]interface{}{"count": commentCount, "data": nil}

	userHasCollect, _ := services.IsUserHasCollect(id, 1, 0)

	specification, _ := services.GetSpecificationListByGoodsId(good.BrandId)
	product, _ := services.GetProductsByGoodsId(good.BrandId)

	services.SaveFootprint(id, 1)

	data := map[string]interface{}{"info": good, "gallery": gallery, "attribute": attribute,
		"issue": issue, "brand": brand, "comment": comment, "userHasCollect": userHasCollect,
		"specificationList": specification, "productList": product}
	c.Success(&data)
}

func (c *GoodsController) GetRelated() {
	id, err := c.Ctx.URLParamInt("id")
	data, err := services.GetRelatedGoodsById(id)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(data)
	}
}