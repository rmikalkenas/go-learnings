package main

import (
	"errors"
	"fmt"
)

type Thing struct {
	title string
}

type Person struct {
	fullName string
	things   []Thing
}

type Persons struct {
	persons []*Person
}

func (p Persons) FindByFullName(n string) (*Person, error) {
	for _, v := range p.persons {
		if v.fullName == n {
			return v, nil
		}
	}

	return &Person{}, errors.New("Not found")
}

func (p *Person) AddThing(t Thing) {
	p.things = append(p.things, t)
}

func main() {
	tv := Thing{"TV"}
	pc := Thing{"PC"}

	p1 := Person{fullName: "Mary"}
	p2 := Person{fullName: "James"}

	p2.AddThing(pc)

	fmt.Println(p2)

	pp := Persons{persons: []*Person{&p1, &p2}}

	p, _ := pp.FindByFullName("Mary")
	p.AddThing(tv)

	p, _ = pp.FindByFullName("James")
	p.AddThing(tv)

	for _, v := range pp.persons {
		fmt.Println(v.fullName, v.things)
	}
}
