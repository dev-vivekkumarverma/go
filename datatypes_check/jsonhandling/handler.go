package jsonhandling

import (
	"encoding/json"
	"fmt"
)

// Todo: Locks for mutithreading and resource sharing

type Person struct {
	Naam  string `json:"Name,omitempty"`
	Umar  int    `json:"Age"`
	Phone string `json:"Phone_number,omitempty"`
}

func JsonToStruct() Person {

	jsonString := `{"Name":"Bhim Rao Chauhan", "Age":111,"Phone_number":"301020004334"}`

	fmt.Printf("value : %v and type: %T", jsonString, jsonString)

	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("ERROR::", err)
	}

	return person
}

func StructToJson(p Person) string {
	jsonString, err := json.Marshal(p)
	if err != nil {
		fmt.Println("ERROR::", err)
	}
	return string(jsonString)

}
