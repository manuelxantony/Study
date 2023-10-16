package main

import (
	"encoding/json"
	"fmt"
)

type Songs struct {
	Name 	string
	Band 	string
	Album 	string
	Year 	int
}

func main() {
	fruits := map[string]int{
		"apple" : 1,
		"orange": 2,
		"mango": 2,
	}
	bytes, _ := json.Marshal(fruits)
	fmt.Println(string(bytes))

	song := Songs{
		Name: "Spit it out",
		Band: "Slipknot",
		Album: "iowa",
		Year: 2001,
	}

	bytes, _ = json.Marshal(song)
	fmt.Println(string(bytes))

}