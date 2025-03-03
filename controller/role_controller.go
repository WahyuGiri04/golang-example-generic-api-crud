package controller

import (
	baseController "golang-example-generic-api-crud/controller/base"
	"golang-example-generic-api-crud/model"
	baseModel "golang-example-generic-api-crud/model/base"
	"golang-example-generic-api-crud/service"
	"golang-example-generic-api-crud/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	Service *service.RoleService
	baseController.BaseController[model.Role]
}

func NewRoleController() *RoleController {
	roleService := service.NewRoleService()
	return &RoleController{
		Service:        roleService,
		BaseController: baseController.BaseController[model.Role]{Service: roleService.BaseService},
	}
}

func (ctrl *RoleController) GetByName(c *gin.Context) {
	name := c.Param("name")
	role, err := ctrl.Service.GetByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, baseModel.BaseResponse{
			Status:  util.Failed,
			Message: "Role not found",
		})
		return
	}

	c.JSON(http.StatusOK, baseModel.BaseResponse{
		Status:  util.Success,
		Message: "Success get role",
		Data:    role,
	})
}
