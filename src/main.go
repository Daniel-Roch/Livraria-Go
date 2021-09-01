package main

import "fmt"

func main() {
	nome := "Daniel" /* declarando variável */
	fmt.Println("Olá Sr.", nome)

	var comando string
	fmt.Scan(&comando)
	fmt.Println("O número foi:", comando)
}
