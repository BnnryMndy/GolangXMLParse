package xmlparse

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Project struct {
	XMLName     xml.Name   `xml:"offer"`
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
	xmlFile, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(xmlFile)

	return data
}

func main() {

}
