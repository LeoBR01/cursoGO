package main

import "fmt"

func main(){
	var variavel1 string = "Variavel1" //Declaracao direta
	variavel2 := "Variavel2" //Declaracão implicita
	fmt.Println(variavel1,variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)
	fmt.Println(variavel3,variavel4)

	variavel5, variavel6 := "Variavel5", "Variavel6"
	fmt.Println(variavel5,variavel6)

	const constante1 string = "Constante 1" //inváriavel
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}