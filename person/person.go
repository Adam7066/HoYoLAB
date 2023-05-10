package person

import (
	"io"
	"os"

	"github.com/Adam7066/golang/log"
)

type Person struct {
	Name   string `json:"name"`
	Ltuid  string `json:"ltuid"`
	Ltoken string `json:"ltoken"`
}

type Persons []Person

func GetPersons() string {
	f, err := os.Open("persons.json")
	if err != nil {
		log.Error.Println(err)
		return ""
	}
	defer f.Close()
	var chunk []byte
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Error.Println(err)
			return ""
		}
		if n == 0 {
			break
		}
		chunk = append(chunk, buf[:n]...)
	}
	return string(chunk)
}
