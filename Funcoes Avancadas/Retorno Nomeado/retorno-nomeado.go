package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2      // Nao precisa por os : para declarar a variavel pois ja esta nomeada no retorno da funcao
    subtracao = n1 - n2
    return
}


func main()  {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}