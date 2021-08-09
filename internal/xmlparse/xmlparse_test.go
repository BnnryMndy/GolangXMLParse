package xmlparse

import (
	"testing"
)

func TestGetStructsFromXML(t *testing.T) {
	//Arrange
	var testStruct Project

	testStruct.ID = 109805
	testStruct.Name = "ЛЕНИНГРАDКА 58"
	testStruct.Address = "Ленинградское шоссе, 58"

	//Act
	testProject := getStructsFromXML(testData)

	//Assert
	if testProject[0].ID != testStruct.ID {
		t.Error("Expected ", testStruct.ID, ", got ", testProject[0].ID)
	}

	if testProject[0].Name != testStruct.Name {
		t.Error("Expected ", testStruct.Name, ", got ", testProject[0].Name)
	}

	if testProject[0].Address != testStruct.Address {
		t.Error("Expected ", testStruct.Address, ", got ", testProject[0].Address)
	}
}


var testData = []byte(`<offer internal-id="109805">
<type>продажа</type>
<property-type>жилая</property-type>
<category>квартира</category>
<deal-status>первичная продажа</deal-status>
<creation-date>2021-08-05T07:44:02+03:00</creation-date>
<last-update-date>2021-08-05T07:44:02+03:00</last-update-date>
<location>
  <country>Россия</country>
  <region>Москва</region>
  <district>Головинский  район</district>
  <address>Ленинградское шоссе, 58</address>
  <latitude>55.843358</latitude>
  <longitude>37.48189</longitude>
  <metro>
	<name>Водный стадион</name>
  </metro>
  <metro>
	<name>Речной вокзал</name>
  </metro>
</location>
<sales-agent>
  <organization>Группа ЛСР</organization>
  <phone>+7 (495) 228-22-88</phone>
  <category>застройщик</category>
</sales-agent>
<price>
  <value>21297600</value>
  <currency>RUB</currency>
</price>
<mortgage>да</mortgage>
<area>
  <value>69.60</value>
  <unit>кв. м</unit>
</area>
<living-space>
  <value>33.20</value>
  <unit>кв. м</unit>
</living-space>
<kitchen-space>
  <value>22.50</value>
  <unit>кв. м</unit>
</kitchen-space>
<rooms>2</rooms>
<new-flat>1</new-flat>
<floor>2</floor>
<floors-total>29</floors-total>
<building-name>ЛЕНИНГРАDКА 58</building-name>
<built-year>2021</built-year>
<building-state>unfinished</building-state>
<ready-quarter>3</ready-quarter>
<building-section>2</building-section>
<description>В продаже 2-комнатная евро квартира  в новом доме жилого комплекса "ЛЕНИНГРАDКА 58". Квартира в корпусе "Ленинградское шоссе" на 2 этаже, общая площадь двухкомнатной квартиры 69.6 м.кв., кухня 22.5 м.кв. Жилой комплекс "ЛЕНИНГРАDКА 58" располагается по адресу Ленинградское шоссе, 58 (Головинский  район г. Москва), рядом станция метро Водный стадион. Срок сдачи новостройки - 2021. . Стоимость квартиры - 21 297 600 рублей. Для вашего удобства предусмотрено онлайн бронирование на сайте Застройщика. Легко. Быстро. Бесплатно.</description>
<yandex-building-id>1570842</yandex-building-id>
<yandex-house-id>1570893</yandex-house-id>
<building-type>монолит</building-type>
</offer>`)
