package service

import (
	"github.com/BnnryMndy/GolangXMLParse/internal/repository"
	"github.com/BnnryMndy/GolangXMLParse/internal/xmlparse"
	"github.com/gin-gonic/gin"
)

type parseService struct {
	repo repository.Repository
}

func NewParseService(repo *repository.Repository) *parseService {
	return &parseService{
		repo: *repo,
	}
}

func (p *parseService) Parse(c *gin.Context) ([]xmlparse.Project, error) {

	//TODO: найти элегантное решение без использования перебора и доп. полей в структурах
	var lot []xmlparse.Lot
	err := c.BindXML(&lot)
	if err != nil {
		return nil, err
	}

	var section []xmlparse.Section
	err = c.BindXML(&section)
	if err != nil {
		return nil, err
	}

	for _, lel := range lot {
		for j, sel := range section {
			if lel.ID == sel.ParrentID {
				section[j].Lot = append(section[j].Lot, lel)
			}
		}
	}

	var building []xmlparse.Building
	err = c.BindXML(&building)
	if err != nil {
		return nil, err
	}

	for _, sel := range section {
		for j, bel := range building {
			if sel.ParrentID == bel.ParrentID {
				building[j].Section = append(building[j].Section, sel)
			}
		}
	}

	var project []xmlparse.Project
	err = c.BindXML(&project)
	if err != nil {
		return nil, err
	}

	for _, bel := range building {
		for j, el := range project {
			if bel.ParrentID == el.ID {
				project[j].Building = append(project[j].Building, bel)
			}
		}
	}

	return project, nil
}
