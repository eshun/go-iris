package api

import (
	"demo/controllers"
	"demo/services"
)

type IndexController struct{
	controllers.ApiController
}

func (c *IndexController) Get() {
	banner, _ := services.GetAdsByPositionId(1)
	channel, _ := services.GetChannels()
	newGoods, _ := services.GetNewGoods()
	hotGoods, _ := services.GetHotGoods()
	brand, _ := services.GetBrands()
	topic, _ := services.GetTopics(3)
	category, _ := services.GetCategorys()

	data := &map[string]interface{}{"banner": banner, "channel": channel, "newGoodsList": newGoods, "hotGoodsList": hotGoods, "brandList": brand, "topicList": topic, "categoryList": category}

	c.Success(data)
}