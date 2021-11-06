package stringornumber

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
)

type Data struct {
	XMLName xml.Name `xml:"data" json:"-"`
	Users []User `xml:"user" json:"user"`
}

type User struct {
	ID      int64   `json:"id" xml:"id"`
	Address Address `json:"address" xml:"address"`
	Age     int     `json:"age" xml:"age"`
}

func (u *User) UnmarshalJSON(data []byte) error {
	var objMap map[string]*json.RawMessage
	if err := json.Unmarshal(data, &objMap); err != nil {
		return err
	}

	if err := json.Unmarshal(*objMap["id"], &u.ID); err != nil {
		return err
	}

	if err := u.Address.UnmarshalJSON(*objMap["address"]); err != nil {
		return err
	}

	if err := json.Unmarshal(*objMap["age"], &u.Age); err != nil {
		var strAge string
		if err = json.Unmarshal(*objMap["age"], &strAge); err != nil {
			return err
		}

		age, err := strconv.Atoi(strAge)
		if err != nil {
			return err
		}

		u.Age = age
	}

	return nil
}

type Address struct {
	CityID int64  `json:"city_id" xml:"city_id"`
	Street string `json:"street" xml:"street"`
}

func (a *Address) UnmarshalJSON(data []byte) error {
	var objMap map[string]*json.RawMessage
	if err := json.Unmarshal(data, &objMap); err != nil {
		return err
	}

	if err := json.Unmarshal(*objMap["city_id"], &a.CityID); err != nil {
		var strCityID string
		err = json.Unmarshal(*objMap["city_id"], &strCityID)
		if err != nil {
			return err
		}
		cityID, err := strconv.Atoi(strCityID)
		if err != nil {
			return err
		}
		a.CityID = int64(cityID)
	}

	if err := json.Unmarshal(*objMap["street"], &a.Street); err != nil {
		return err
	}

	return nil
}

func JsonUnmarshalArrayOfUsers(users *[]User, jsonData []byte) {
	var objMap []*json.RawMessage
	err := json.Unmarshal(jsonData, &objMap)
	if err != nil {
		panic(err)
	}

	for _, v := range objMap {
		user := new(User)
		if err = user.UnmarshalJSON(*v); err != nil {
			panic(err)
		}
		*users = append(*users, *user)
	}
}

func XmlUnmarshalArrayOfUsers(data *Data, rawXml []byte) error {
	if err := xml.Unmarshal(rawXml, data); err != nil {
		return err
	}

	return nil
}
