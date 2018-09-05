package api

import (
	"demo/controllers"
	"demo/services"
	"demo/models"
)

type BrandController struct{
	controllers.ApiController
}

func (c *BrandController) GetList() {
	retPage := models.RetPage{}
	var start int
	size, err := c.Ctx.URLParamInt("size")
	if err != nil || size < 1 {
		size = 10
	}
	retPage.Pagesize = size
	page, err := c.Ctx.URLParamInt("page")
	if err != nil || page < 2 {
		page = 1
		start = 0
	} else {
		start = size*(page-1) + 1
	}
	retPage.CurrentPage = page
	count, err := services.GetTopicCount()
	if int(count) > size {
		retPage.TotalPages = (int(count) + 1) / size
	} else {
		retPage.TotalPages = 1
	}
	list, err := services.GetPageBrands(size, start)
	if err != nil {
		c.Fail(err)
	} else {
		retPage.Data = list
		retPage.Count = len(list)
		c.Success(retPage)
	}
}

func (c *BrandController) GetDetail() {
	id, err := c.Ctx.URLParamInt("id")
	model, err := services.GetBrandById(id)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(&map[string]interface{}{"brand": model})
	}
}