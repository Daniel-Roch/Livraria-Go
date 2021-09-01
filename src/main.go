package main

import "fmt"

func main() {
	nome := "Daniel" /* declarando variável */
	fmt.Println("Olá Sr.", nome)

	fmt.Println("Escolha um número de 1 a 10")
	var comando string
	fmt.Scan(&comando)
	fmt.Println("O número foi:", comando)
}
