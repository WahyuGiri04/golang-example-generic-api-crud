package routes

import (
	"golang-example-generic-api-crud/controller"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine) {
	var path = os.Getenv("CONTEXT_PATH")
	RoleController := controller.NewRoleController()
	roleRoutes := route.Group(path + "/role")
	{
		roleRoutes.POST("/create", RoleController.Create)
		roleRoutes.GET("/get-all", RoleController.GetAll)
		roleRoutes.GET("/get-by-id", RoleController.GetById)
		roleRoutes.PUT("/update", RoleController.Update)
		roleRoutes.DELETE("/delete", RoleController.Delete)
		roleRoutes.GET("/get-pagination", RoleController.GetPagination)
		roleRoutes.GET("/get-by-name", RoleController.GetByName)
	}
}
