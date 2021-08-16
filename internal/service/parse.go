package service

import (
	"github.com/BnnryMndy/GolangXMLParse/internal/repository"
	"github.com/BnnryMndy/GolangXMLParse/internal/xmlparse"
)

type parseService struct {
	repo repository.Repository
}

func NewParseService(repo *repository.Repository) *parseService {
	return &parseService{
		repo: *repo,
	}
}

func (p *parseService) Parse(data []byte) []xmlparse.Project {

	return xmlparse.GetStructsFromXML(data)
}
