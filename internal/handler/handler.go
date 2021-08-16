package handlers

import (
	"github.com/BnnryMndy/GolangXMLParse/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandlers(serv *service.Service) *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		parse := api.Group("/parse")
		{
			//parse.GET("/", h.getParse)
			parse.POST("/", h.parse)
			save := parse.Group("/save")
			{
				//save.GET("/", h.getParseSave)
				save.POST("/", h.postParseSave)
			}
		}
	}

	return router
}
