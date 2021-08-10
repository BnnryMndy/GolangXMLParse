package xmlparse

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

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

	for _, lel := range lot {
		for j, sel := range section {
			if lel.ID == sel.ParrentID {
				section[j].Lot = append(section[j].Lot, lel)
			}
		}
	}

	var building []Building
	xml.Unmarshal(data, &building)

	for _, sel := range section {
		for j, bel := range building {
			if sel.ParrentID == bel.ParrentID {
				building[j].Section = append(building[j].Section, sel)
			}
		}
	}

	var project []Project
	xml.Unmarshal(data, &project)

	for _, bel := range building {
		for j, el := range project {
			if bel.ParrentID == el.ID {
				project[j].Building = append(project[j].Building, bel)
			}
		}
	}

	return project
}
