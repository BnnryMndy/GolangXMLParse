package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// func (h *Handler) getParse(c *gin.Context) {
// }

func (h *Handler) parse(c *gin.Context) {
	data, err := c.GetRawData()

	if err != nil {
		log.Print(fmt.Errorf("Error: %w", err))
	}

	h.services.Parse.Parse(data)
}
