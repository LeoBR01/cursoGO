package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string //itens de array sao sempre do mesmo tipo.
	array1[0] = "posicao 1"

	fmt.Println(array1)

	array2 := [5]string{"posicao 1", "posicao 2", "posicao 3", "posicao 4", "posicao 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} //Nao torna o array dinamico, mas pode ir preenchendo sem especificar tamanho
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17} //itens de slice sao sempre do mesmo tipo.
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice)) // Slice aponta para um array

	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3] //Isso nao altera o array, só faz um 'ponteiro'
	fmt.Println(slice2)

	array2[1] = "Posicao alterada"
	fmt.Println(slice2)

	// Array Interno
	slice3 := make([]float32, 10, 11)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6) //Tentando estourar a capacidade do slice , que a principio nao tem limite
							   //O go chega no limite e duplica o valor da capacidade, criando outro array para referenciar.

	fmt.Println("slice3>:", slice3)
	fmt.Println("slice3_tamanho:", len(slice3))
	fmt.Println("slice3_capacidade:", cap(slice3))

	

}
