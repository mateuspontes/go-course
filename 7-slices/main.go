package main

import "fmt"

func main() {
	linguagens := make([]string, 0, 5)
	linguagens = append(linguagens, "python", "ruby", "javascript")
	linguagens = append(linguagens, "go")

	fmt.Println("Linguagens da globo.com: ", linguagens)
	fmt.Println("Linguagens da globo.com: ", linguagens[1:3])

	for i, linguagem := range linguagens {
		fmt.Printf(
			"Linguagem na posição %d do slice, valor: %s\n",
			i,
			linguagem,
		)
	}
}
