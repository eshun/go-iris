package services

import (
	"demo/models"
	"time"
	"strconv"
)

func GetKeyword(userId int) (*map[string]interface{}) {
	var defaultKeyword *models.Keywords
	models.Orm.Where("is_default = ?", 1).Get(&defaultKeyword)

	var hotList []*models.Keywords
	models.Orm.Distinct("keyword").Cols("keyword,is_hot").Limit(10, 0).Find(&hotList)

	var historyList []*models.SearchHistory
	models.Orm.Distinct("keyword").Cols("keyword").Where("user_id = ?", userId).Limit(10, 0).Find(&historyList)

	return &map[string]interface{}{"defaultKeyword": defaultKeyword, "hotKeywordList": hotList, "historyKeywordList": historyList}
}

func ClearSearchHistoryByUserId(userId int) (int64, error) {
	model := new(models.SearchHistory)
	return models.Orm.Where("user_id = ?", userId).Delete(&model)
}

func GetHelperByKeyword(keyword string) (*map[string]interface{}) {
	var list []*models.Keywords
	models.Orm.Distinct("keyword").Cols("keyword").Where("keyword like ?", keyword+"%").Find(&list)
	var data = make(map[string]interface{})

	if len(list) > 0 {
		for _, v := range list {
			data["keyword"] = v.Keyword
		}
	}
	return &data
}

func SaveSearchHistory(userId int,keyword string) (*models.SearchHistory,error) {
	model := models.SearchHistory{Keyword: keyword, UserId: strconv.Itoa(userId), AddTime: int(time.Now().Unix())}
	_, err := models.Orm.Insert(&model)
	return &model, err
}