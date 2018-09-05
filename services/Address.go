package services

import "demo/models"

func GetAddressByUserId(userId int) ([]*models.Address, error) {
	var list []*models.Address
	err := models.Orm.Where("user_id = ?", userId).Find(&list)
	if len(list) > 0 {
		for i, v := range list {
			country, _ := GetRegionById(v.CountryId)
			if country != nil {
				list[i].CountryName = country.Name
				list[i].FullRegion += country.Name
			}
			province, _ := GetRegionById(v.ProvinceId)
			if country != nil {
				list[i].ProvinceName = province.Name
				list[i].FullRegion += province.Name
			}
			city, _ := GetRegionById(v.CityId)
			if country != nil {
				list[i].CityName = city.Name
				list[i].FullRegion += city.Name
			}
		}
	}
	return list, err
}

func GetAddressById(id int) (*models.Address, error) {
	var model *models.Address
	has, err := models.Orm.Id(id).Get(&model)
	if has {
		country, _ := GetRegionById(model.CountryId)
		if country != nil {
			model.CountryName = country.Name
			model.FullRegion += country.Name
		}
		province, _ := GetRegionById(model.ProvinceId)
		if country != nil {
			model.ProvinceName = province.Name
			model.FullRegion += province.Name
		}
		city, _ := GetRegionById(model.CityId)
		if country != nil {
			model.CityName = city.Name
			model.FullRegion += city.Name
		}
	}
	return model, err
}

func DeleAddressById(id int) (int64, error) {
	return models.Orm.Id(id).Delete(&models.Address{})
}

func SaveAddress(address *models.Address) (int64, error) {
	if (address.IsDefault == 1) {
		models.Orm.Where("user_id = ?", address.UserId).Cols("is_default").Update(&models.Address{IsDefault: 0})
	}
	has, _ := models.Orm.Get(&address)
	if has {
		return models.Orm.Update(&address)
	} else {
		return models.Orm.Insert(&address)
	}
}