package main

import "fmt"

func main() {
	people := map[string]float64{
		"wilson": 25,
		"andre":  26,
	}

	people["junior"] = 10

	for name, age := range people {
		fmt.Printf("Chave %s do mapa valor: %v\n", name, age)
	}

	delete(people, "wilson")
	fmt.Printf("Posição não alocada: %v\n", people["wilson"])
}
