package model

import "golang-example-generic-api-crud/model/base"

type Users struct {
	baseModel.BaseModel
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
