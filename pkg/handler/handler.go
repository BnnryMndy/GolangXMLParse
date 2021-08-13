package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		parse := router.Group("/parse")
		{
			parse.GET("/", h.getParse)
			parse.POST("/", h.postParse)
			save := router.Group("/save")
			{
				parse.GET("/")
				parse.POST("/")
			}
		}
	}

	return router
}
