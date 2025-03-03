package model

import "golang-example-generic-api-crud/model/base"


type Role struct {
	baseModel.BaseModel
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (Role) TableName() string {
	return "ROLE"
}
