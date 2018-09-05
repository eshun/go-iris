package api

import (
	"demo/controllers"
	"demo/services"
	"demo/models"
)

type AddressController struct{
	controllers.ApiController
}

func (c *AddressController) GetBy(id int) {
	data, err := services.GetAddressById(id)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(data)
	}
}

func (c *AddressController) GetList() {
	data, err := services.GetAddressByUserId(1)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(data)
	}
}

func (c *AddressController) PostSaveBy(id ,province_id,city_id,district_id,user_id,is_default int,name,mobile,address string) {
	model := models.Address{Id: id, Name: name, Mobile: mobile, Address: address, ProvinceId: province_id, CityId: city_id, DistrictId: district_id, UserId: user_id, IsDefault: is_default}
	_, err := services.SaveAddress(&model)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success(model)
	}
}

func (c *AddressController) PostDeleBy(id int) {
	_, err := services.DeleAddressById(id)
	if err != nil {
		c.Fail(err)
	} else {
		c.Success("")
	}
}