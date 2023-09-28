package main

import "fmt"

func somar(n1 int8,n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1,n2 int8) (int8, int8){ // Se os param tiverem o mesmo tipo, podemos declarar so o ultimo
	soma:= n1+n2								   // A funcao pode ter mais de um retorno
	subtracao := n1-n2
	return soma,subtracao
} 

func main(){
	soma := somar(10,20)
	fmt.Println(soma)

	var f = func(txt string) string{
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da funcao 1")
	fmt.Println(resultado)

	resultadosSoma, resultadosSubtracao := calculosMatematicos(10,15) // Devemos criar uma variavel para cada retorno
	fmt.Println(resultadosSoma,resultadosSubtracao) //Se eu n quiser usar o valor da variavel , podemos usar um `_` para o Go ignorar a variavel
	


}
