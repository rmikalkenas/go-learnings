package main

import "fmt"

type person struct {
	first           string
	last            string
	favoriteFlavors []string
}

func main() {
	p1 := person{
		first:           "Rokas",
		last:            "Mikalkenas",
		favoriteFlavors: []string{"Chocolate", "Vanilla"},
	}

	p2 := person{
		first:           "Jolita",
		last:            "Gubeviciute",
		favoriteFlavors: []string{"With nuts", "Tiramisu", "Mint"},
	}

	persons := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(persons)

	fmt.Println("----------")

	p1.printPerson()
	fmt.Println("----------")
	p2.printPerson()

	anon := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Anony",
		friends: map[string]int{
			"A1": 1,
			"A2": 93,
		},
		favDrinks: []string{"Cola", "Fanta"},
	}

	fmt.Println(anon)
}

func (p person) printPerson() {
	fmt.Printf("First name: %v\n", p.first)
	fmt.Printf("Last name: %v\n", p.last)
	fmt.Println("Favorite ice cream flavors:")
	for _, flavor := range p.favoriteFlavors {
		fmt.Printf("\t%v\n", flavor)
	}
}
