package routes

import (
	"golang-example-generic-api-crud/controller"
	"golang-example-generic-api-crud/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine) {
	var path = os.Getenv("CONTEXT_PATH")

	RoleController := controller.NewRoleController()
	// Endpoint health check
	route.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	roleRoutes := route.Group(path + "/role")
	roleRoutes.Use(middleware.AuthMiddleware())
	{
		roleRoutes.POST("/create", RoleController.Create)
		roleRoutes.GET("/get-all", RoleController.GetAll)
		roleRoutes.GET("/get-by-id", RoleController.GetById)
		roleRoutes.PUT("/update", RoleController.Update)
		roleRoutes.DELETE("/delete", RoleController.Delete)
		roleRoutes.GET("/get-pagination", RoleController.GetPagination)
		roleRoutes.GET("/get-by-field", RoleController.GetByField)
		roleRoutes.GET("/get-by-name", RoleController.FindByName)
	}
}
