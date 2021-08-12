package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		parse := router.Group("/parse")
		{
			parse.GET("/")
			parse.POST("/")
			save := router.Group("/save")
			{

			}
		}
	}

	return router
}
