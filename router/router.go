package router

import (
	"github.com/gin-gonic/gin"

	"github.com/n1ce37/aws/api"
	"github.com/n1ce37/aws/api/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	baseRouter := r.Group("/api")

	userRouter := baseRouter.Group("/users")
	{
		userRouter.GET("/", api.ListUsers)
		userRouter.POST("/uuid", api.CreateUUID)
	}

	templateRouter := baseRouter.Group("/templates")
	{
		templateRouter.POST("/", api.CreateTemplate)
		templateRouter.GET("/", api.ListTemplates)
		templateRouter.GET("/:id", api.ReadTemplate)
		templateRouter.PATCH("/:id", api.UpdateTemplate)
		templateRouter.DELETE("/:id", api.DeleteTemplate)
	}

	historyRouter := baseRouter.Group("/histories", middleware.GetUUID)
	{
		historyRouter.POST("/", api.CreateHistory)
		historyRouter.GET("/", api.ListHistories)
		historyRouter.GET("/:id", api.ReadHistory)
		historyRouter.PATCH("/:id", api.UpdateHistory)
		historyRouter.DELETE("/:id", api.DeleteHistory)
	}

	return r
}
