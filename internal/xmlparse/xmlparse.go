package xmlparse

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Project struct {
	ID          int32      `xml:"internal-id,attr"`
	Name        string     `xml:"building-name"`
	Description string     `xml:"description"`
	Address     string     `xml:"location>address"`
	Building    []Building `xml:"offer"`
}

type Building struct {
	ID      int32 `xml:"yandex-house-id"`
	Name    string
	Section []Section `xml:"offer"`
}

type Section struct {
	ID   int32 `xml:"yandex-building-id"`
	Name string
	Lot  []Lot `xml:"offer"`
}

type Lot struct {
	ID            int32
	Floor         int32  `xml:"floor"`
	TotalSquare   int32  `xml:"area>value"`
	LivingSquare  int32  `xml:"living-space>value"`
	KitchenSquare int32  `xml:"kitchen-space>value"`
	Price         int32  `xml:"price>value"`
	LotType       string `xml:"type"`
	RoomType      string `xml:"category"`
}

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

func getStructsFromXML(data []byte) []Project {
	var project []Project
	xml.Unmarshal(data, &project)
	return project
}
