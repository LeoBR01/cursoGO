package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// if_init
	if outroNumero := numero; outroNumero > 0 {       // Ao criar uma variavel dentro de um if, ela fica sujeita aquele escopo do if. Nao é possivel acessar de fora.
		fmt.Println("Numero é maior que zero")
	} else {
		fmt.Println("Numero é menor ou igual a zero")
	}
}
