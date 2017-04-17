package main

import "fmt"

// Person struct
type Person struct {
	Name   string
	Age    int
	Active bool
}

func changePerson(p Person) {
	p.Age = 90
}

func main() {
	p := Person{
		Name:   "Mateus Pontes",
		Age:    29,
		Active: true,
	}
	changePerson(p)
	fmt.Println("Nome", p.Name)
	fmt.Println("Age", p.Age)
	fmt.Println("Active", p.Active)

	p2 := &Person{Name: "Joao"}
	fmt.Println("Nome", p2.Name)
	fmt.Println("Age", p2.Age)
	fmt.Println("Active", p2.Active)
}
