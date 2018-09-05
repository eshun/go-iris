package api

import (
	"demo/controllers"
	"demo/services"
)

type RegionController struct{
	controllers.ApiController
}

func (c *RegionController) GetInfoBy(id int) {
	model, err := services.GetRegionById(id)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(model)
	}
}

func (c *RegionController) GetListBy(parentId int) {
	list, err := services.GetRegionByParentId(parentId)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(list)
	}
}