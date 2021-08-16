package service

import (
	"github.com/BnnryMndy/GolangXMLParse/internal/repository"
	"github.com/BnnryMndy/GolangXMLParse/internal/xmlparse"
)

type Parse interface {
	Parse(data []byte) []xmlparse.Project
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
