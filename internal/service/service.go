package service

import (
	"github.com/BnnryMndy/GolangXMLParse/internal/repository"
	"github.com/BnnryMndy/GolangXMLParse/internal/xmlparse"
	"github.com/gin-gonic/gin"
)

type Parse interface {
	Parse(c *gin.Context) ([]xmlparse.Project, error)
}

type ParseSave interface {
}

type Service struct {
	Parse
	ParseSave
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		Parse: NewParseService(repos),
	}
}
