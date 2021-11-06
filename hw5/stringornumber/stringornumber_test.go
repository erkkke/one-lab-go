package stringornumber

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	rawJson = []byte(`[
	  {
		"id": 1,
		"address": {
		  "city_id": 5,
		  "street": "Satbayev"
		},
		"age": 20
	  },
	  {
		"id": 1,
		"address": {
		  "city_id": "6",
		  "street": "Al-Farabi"
		},
		"age": "32"
	  }
	]`)

	rawXml = []byte(`
	<data>
		<user>
			<id>1</id>
			<address>
				<city_id>5</city_id>
				<street>Satbayev</street>
			</address>
			<age>20</age>
		</user>
		<user>
			<id>1</id>
			<address>
				<city_id>6</city_id>
				<street>Al-Farabi</street>
			</address>
			<age>32</age>
		</user>
	</data>`)
)

func TestJsonUnmarshalArrayOfUsers(t *testing.T) {
	var result Data
	JsonUnmarshalArrayOfUsers(&result.Users, rawJson)

	expected := Data {
		Users: []User{
			{
				ID: 1,
				Address: Address{
				CityID: 5,
				Street: "Satbayev"},
				Age: 20,
			},
			{
				ID: 1,
				Address: Address{
					CityID: 6,
					Street: "Al-Farabi",
				},
				Age: 32,
			},
		},
	}

	assert.Equal(t, expected.Users, result.Users)
}

func TestXmlUnmarshalArrayOfUsers(t *testing.T) {
	var result Data
	if err := XmlUnmarshalArrayOfUsers(&result, rawXml); err != nil {
		_ = fmt.Errorf("error: %v", err)
	}

	expected := Data {
		Users: []User{
			{
				ID: 1,
				Address: Address{
					CityID: 5,
					Street: "Satbayev"},
				Age: 20,
			},
			{
				ID: 1,
				Address: Address{
					CityID: 6,
					Street: "Al-Farabi",
				},
				Age: 32,
			},
		},
	}

	assert.Equal(t, expected.Users, result.Users)
}


