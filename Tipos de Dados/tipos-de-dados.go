package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	//uint => unsygned int (int sem sinal)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 12345
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.34 //sempre usando ponto, virgula nao funciona
	fmt.Println(numeroReal1)

	numeroReal2 := 1234.54 //inferencia de tipos, não existe apenas o float
	fmt.Println(numeroReal2)

	char := "A" // print a valor do caracter na tabela ASCI
	fmt.Println(char)

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
