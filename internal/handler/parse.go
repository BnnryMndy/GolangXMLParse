package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BnnryMndy/GolangXMLParse/internal/xmlparse"
	"github.com/gin-gonic/gin"
)

func (h *Handler) parse(c *gin.Context) {

	var bindedProject xmlparse.Project

	errXML := c.BindXML(&bindedProject)
	if errXML != nil {
		log.Print(fmt.Errorf("Error: %w", errXML))
	}
	log.Printf("project id: %d", bindedProject.ID)

	data, err := c.GetRawData()

	if err != nil {
		log.Print(fmt.Errorf("Error: %w", err))
	}

	if data == nil {
		log.Printf("data is nil!")
	}

	result := h.services.Parse.Parse(data)

	if result != nil {
		log.Print("result was reached")
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "handler work, may be",
	})
}
