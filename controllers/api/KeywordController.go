package api

import (
	"demo/controllers"
	"demo/services"
)

type KeywordController struct{
	controllers.ApiController
}

func (c *KeywordController) Get() {
	data := services.GetKeyword(1)
	c.Success(data)
}

func (c *KeywordController) GetHelper(keyword string) {
	data := services.GetHelperByKeyword(keyword)
	c.Success(data)
}

func (c *KeywordController) PostClearHistory() {
	_, err := services.ClearSearchHistoryByUserId(1)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success("")
	}
}