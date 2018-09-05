package controllers

import (
	"encoding/json"
	"demo/models"
)

type UserController struct {}

// GET: http://localhost:8080/user
func (m *UserController) Get() string {
	result, _ := json.Marshal(&models.RetMsg{})
	return string(result)
}

// GET: http://localhost:8080/user/profile/followers/{id:long}
func (m *UserController) GetProfileFollowers(id int64) string { return "MyCustomHandler says Hey" }