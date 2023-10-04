package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parametro")

	fmt.Println(retorno)
} 