package xmlparse

type Project struct {
	ID          int32  `xml:"internal-id,attr"`
	Name        string `xml:"building-name"` //? Не уверен, что записывать сюда
	Description string `xml:"description"`
	Address     string `xml:"location>address"`
	Building    []Building
}

type Building struct {
	ParrentID int32 `xml:"internal-id,attr"` //TODO: найти элегантное решение без использования перебора
	ID        int32 `xml:"yandex-house-id"`
	Name      string
	Section   []Section
}

type Section struct {
	ParrentID int32 `xml:"internal-id,attr"` //TODO: найти элегантное решение без использования перебора
	ID        int32 `xml:"yandex-building-id"`
	Name      string
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
