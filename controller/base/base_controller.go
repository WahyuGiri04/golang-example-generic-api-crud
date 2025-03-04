package baseController

import (
	baseModel "golang-example-generic-api-crud/model/base"
	baseService "golang-example-generic-api-crud/service/base"
	"golang-example-generic-api-crud/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BaseController[T any] struct {
	Service *baseService.BaseService[T]
}

func (ctrl *BaseController[T]) Create(c *gin.Context) {
	var entity T

	if !util.BindJSONGeneric(c, &entity) {
		return
	}

	if err := ctrl.Service.Create((&entity)); err != nil {
		c.JSON(http.StatusInternalServerError, baseModel.BaseResponse{
			Status:  util.Failed,
			Message: "Failed to create entity",
		})
		return
	}
	c.JSON(http.StatusOK, baseModel.BaseResponse{
		Status:  util.Success,
		Message: "Success to create entity",
		Data:    entity,
	})
}

func (ctrl *BaseController[T]) GetAll(c *gin.Context) {
	var entities []T
	if err := ctrl.Service.GetAll(&entities); err != nil {
		c.JSON(http.StatusInternalServerError, baseModel.BaseResponse{
			Status:  util.Failed,
			Message: "Failed to get all entity",
		})
		return
	}
	c.JSON(http.StatusOK, baseModel.BaseResponse{
		Status:  util.Success,
		Message: "Success to get all entity",
		Data:    entities,
	})
}

func (ctrl *BaseController[T]) GetById(c *gin.Context) {
	var entity T
	id, _ := strconv.Atoi(c.Query("id"))
	if err := ctrl.Service.GetById(uint(id), &entity); err != nil {
		c.JSON(http.StatusNotFound, baseModel.BaseResponse{
			Status:  util.Failed,
			Message: "Failed to get entity by id",
		})
	}
	c.JSON(http.StatusOK, baseModel.BaseResponse{
		Status:  util.Success,
		Message: "Success to get entity by id",
		Data:    entity,
	})
}

func (ctrl *BaseController[T]) Update(c *gin.Context) {
	var entity T
	id, _ := strconv.Atoi(c.Query("id"))
	if !util.BindJSONGeneric(c, &entity) {
		return
	}

	if err := ctrl.Service.Update(uint(id), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, baseModel.BaseResponse{
			Status:  util.Failed,
			Message: "Failed to update entity",
		})
		return
	}
	c.JSON(http.StatusOK, baseModel.BaseResponse{
		Status:  util.Success,
		Message: "Success to update entity",
		Data:    entity,
	})
}

func (ctrl *BaseController[T]) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, baseModel.BaseResponse{Status: util.Failed, Message: "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, baseModel.BaseResponse{Status: util.Success, Message: "Deleted successfully"})
}

func (ctrl *BaseController[T]) GetPagination(c *gin.Context) {
	var entities []T

	// Ambil parameter page dan pageSize
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	pagination, err := ctrl.Service.GetPagination(page, pageSize, &entities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, baseModel.BaseResponse{
			Status:  util.Failed,
			Message: "Failed to get paginated data",
		})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, baseModel.BaseResponse{
		Status:  util.Success,
		Message: "Success get paginated data",
		Data:    pagination,
	})
}

func (ctrl *BaseController[T]) GetByField(c *gin.Context) {
	field := c.DefaultQuery("field", "name") // Default cari berdasarkan "name"
	value := c.Query("value")

	results, err := ctrl.Service.GetByField(field, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, baseModel.BaseResponse{
			Status:  util.Failed,
			Message: "Failed to get data",
		})
		return
	}

	c.JSON(http.StatusOK, baseModel.BaseResponse{
		Status:  util.Success,
		Message: "Success get data",
		Data:    results,
	})
}

func (ctrl *BaseController[T]) FindByName(c *gin.Context) {
	name := c.Query("name")
	results, err := ctrl.Service.FindByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, baseModel.BaseResponse{
			Status:  util.Failed,
			Message: "Failed to get data",
		})
		return
	}

	c.JSON(http.StatusOK, baseModel.BaseResponse{
		Status:  util.Success,
		Message: "Success get data",
		Data:    results,
	})
}
