package routes

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	api := route.Group("/cgi/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/data-collection", data_collection)
		}
	}
}
