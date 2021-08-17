package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) parse(c *gin.Context) {

	_, err := h.services.Parse.Parse(c)

	if err != nil {
		log.Print(fmt.Errorf("Error while parse: %w", err))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "handler work, may be",
	})
}
