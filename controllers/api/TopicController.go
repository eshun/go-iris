package api

import (
	"demo/controllers"
	"demo/services"
	"demo/models"
)

type TopicController struct{
	controllers.ApiController
}

func (c *TopicController) GetList() {
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
	list, err := services.GetPageTopics(size, start)
	if err != nil {
		c.Fail(err)
	} else {
		retPage.Data = list
		retPage.Count = len(list)
		c.Success(retPage)
	}
}

func (c *TopicController) GetDetail() {
	id, err := c.Ctx.URLParamInt("id")
	model, err := services.GetTopicById(id)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(model)
	}
}

func (c *TopicController) GetRelated() {
	id, err := c.Ctx.URLParamInt("id")
	model, err := services.GetTopicById(id)
	if err != nil {
	} else {
		list, _ := services.GetTopicsByCategoryId(model.TopicCategoryId, 4)
		c.Success(list)
	}
}