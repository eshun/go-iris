package api

import (
	"demo/controllers"
	"demo/services"
	"demo/models"
)

type CategoryController struct{
	controllers.ApiController
}

func (c *CategoryController) Get() {
	list, err := services.GetCategoryLimit(10)
	if err != nil {
		c.Fail(err)
	} else {
		id, _ := c.Ctx.URLParamInt("id")
		var model *models.Category
		if id > 0 {
			model, _ = services.GetCategoryAndChildById(id)
		}
		if model == nil {
			model = list[0]
		}
		data := &map[string]interface{}{"categoryList": list, "currentCategory": model}
		c.Success(data)
	}
}

func (c *CategoryController) GetCurrent() {
	id, err := c.Ctx.URLParamInt("id")
	model, err := services.GetCategoryAndChildById(id)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(&map[string]interface{}{"currentCategory": model})
	}
}