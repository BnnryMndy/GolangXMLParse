package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		parse := api.Group("/parse")
		{
			parse.GET("/", h.getParse)
			parse.POST("/", h.postParse)
			save := parse.Group("/save")
			{
				save.GET("/", h.getParseSave)
				save.POST("/", h.postParseSave)
			}
		}
	}

	return router
}
