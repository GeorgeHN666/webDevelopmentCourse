package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name  string   `json:"Name"`
	Last  string   `json:"Last"`
	Stuff []string `json:"Stuff"`
}

var data person

func main() {
	s := `{"Name":"james","Last":"bond","Stuff":["suit","Gun","Black Sense Of Humor"]}`

	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(s)

	for i, v := range data.Stuff {
		fmt.Println(i, v)
	}

	fmt.Println(data.Name)
	fmt.Println(data.Last)

}
