package main

import "fmt"

func soma(numeros ...int) int { // ...int -> vai receber de 1 a n ints
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {    // Não podemos ter mais de uma parametro variatico por funcao
	for _, numero := range numeros {			 // O parametro variatico tem que obrigatoriamente ser o ultimo parametro da funcao
		fmt.Println(texto, numero)
	}
}

func main() {
	totaldaSoma := soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	fmt.Println(totaldaSoma)

	escrever("Ola mundo", 10, 2, 3, 4, 5, 6, 7)
}
