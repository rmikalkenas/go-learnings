package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	p1 := person{
		First: "Rokas",
		Last:  "Mikalkenas",
		Age:   33,
	}

	persons := []person{p1}

	jb, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1, string(jb))

	pp := []person{}

	err = json.Unmarshal(jb, &pp)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", pp)
}
