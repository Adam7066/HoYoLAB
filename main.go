package main

import (
	"encoding/json"
	"hoyolab/checkin"
	"hoyolab/person"
)

func main() {
	personsData := []byte(person.GetPersons())
	var persons person.Persons
	err := json.Unmarshal(personsData, &persons)
	if err != nil {
		panic(err)
	}
	checkin.Checkin(persons)
}
