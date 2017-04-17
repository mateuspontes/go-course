package main

import "fmt"

// Add - exemplo de função publica
func Add(x, y int) (int, int) {
	return x + y, 0
}

func main() {
	x, y := 10, 20
	result, other := Add(x, y)
	fmt.Printf("O resultado de %d + %d é %d \n", x, y, result)
	fmt.Println("O parâmetro extra é ", other)
}
