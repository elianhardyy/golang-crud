package routers

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/users", controllers.GetAll)
		api.POST("/create", controllers.Add)
		api.PUT("/edit/:id", controllers.Edit)
		api.DELETE("/delete/:id", controllers.Delete)
	}
	return r
}
