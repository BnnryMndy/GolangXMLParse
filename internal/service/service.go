package service

import (
	"github.com/BnnryMndy/GolangXMLParse/internal/repository"
)

type Parse interface {
}

type ParseSave interface {
}

type Service struct {
	Parse
	ParseSave
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{}
}
