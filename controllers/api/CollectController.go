package api

import (
	"demo/controllers"
	"demo/services"
	"demo/models"
)

type CollectController struct{
	controllers.ApiController
}

func (c *CollectController) GetListBy(typeId int) {
	data, err := services.GetCollectByTypeId(1, typeId)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(data)
	}
}

func (c *CollectController) PostAddOrDeleteBy(typeId,valueId int) {
	data, err := services.SaveCollect(&models.Collect{TypeId: typeId, ValueId: valueId})
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(data)
	}
}