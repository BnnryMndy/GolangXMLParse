package handler

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
			parse.POST("/", h.parse)
			save := parse.Group("/save")
			{
				save.POST("/", h.ParseSave)
			}
		}
	}

	return router
}
