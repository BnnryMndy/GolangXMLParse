package xmlparse

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Project struct {
	ID          int32  `xml:"internal-id,attr"`
	Name        string `xml:"building-name"` //? Не уверен, что записывать сюда
	Description string `xml:"description"`
	Address     string `xml:"location>address"`
	Building    []Building
}

type Building struct {
	ParrentID int32     `xml:"internal-id,attr"` //!
	ID        int32     `xml:"yandex-house-id"`
	Name      string    `xml:"building-name"` //? Не уверен, что записывать сюда
	Section   []Section //`xml:"offer"`
}

type Section struct {
	ParrentID int32  `xml:"internal-id,attr"` //! временное решение
	ID        int32  `xml:"yandex-building-id"`
	Name      string `xml:"building-name"` //? Не уверен, что записывать сюда
	Lot       []Lot
}

type Lot struct {
	ID            int32   `xml:"internal-id,attr"` //? Не уверен, что записывать сюда
	Floor         int32   `xml:"floor"`
	TotalSquare   float32 `xml:"area>value"`
	LivingSquare  float32 `xml:"living-space>value"`
	KitchenSquare float32 `xml:"kitchen-space>value"`
	Price         float32 `xml:"price>value"`
	LotType       string  `xml:"type"`
	RoomType      string  `xml:"category"`
}

//getting xml file in byte slice from path.
func getXMLFromPath(path string) []byte {
	xmlFile, fileErr := os.Open(path)

	if fileErr != nil {
		fmt.Errorf("Error while open file: %w", fileErr)
	}

	data, byteErr := ioutil.ReadAll(xmlFile)

	if byteErr != nil {
		fmt.Errorf("Error while read file: %w", byteErr)
	}

	return data
}

//getting struct from byte slice xml
func getStructsFromXML(data []byte) []Project {
	var lot []Lot
	xml.Unmarshal(data, &lot)

	var section []Section
	xml.Unmarshal(data, &section)

	var building []Building
	xml.Unmarshal(data, &building)

	var project []Project
	xml.Unmarshal(data, &project)

	//TODO: Сделать менее затратно
	for i, el := range project {
		for j, bel := range building {
			if el.ID == bel.ParrentID {
				project[i].Building = append(project[i].Building, bel)
			}

			for k, sel := range section {
				if bel.ParrentID == sel.ParrentID {
					building[j].Section = append(building[j].Section, sel)
				}
				for _, lel := range lot {
					if lel.ID == sel.ParrentID {
						section[k].Lot = append(section[k].Lot, lel)
					}

				}
			}
		}
	}

	return project
}
