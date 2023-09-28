package main

import "fmt"

func main(){
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Não podemos realizar operacoes com variaveis que tem tipos diferentes
	// até mesmo numeros com bits diferentes, int32 + int8 nao funciona

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	
	//Operadores relacionais sempre retornam booleanos
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//Operadores logicos
	println("----------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Operadores unários => interagem com apenas 1 variavel
	println("----------")
	numero := 10
	numero++
	numero += 15 //adicionar 15
	numero--
	numero -= 20
	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto)
	
}