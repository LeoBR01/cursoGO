package main

import (
	"fmt"
	//"time"
)

func main()  {
	// i := 0
	// for i < 10 {
    //     i++
    //     fmt.Println("Incrementando i ", i)
    // }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j ", j)
    //     time.Sleep(time.Second)
	// }
    
	nome := [5]string{"a", "b", "c", "d", "e"}

	for indices, nomes := range nome {
		fmt.Println(indices, nomes)
	}

	for indices, nomes := range "PALAVRA" {
		fmt.Println(indices, nomes)
	}

	usuario := map[string]string{
		"nome": "Leonardo",
        "sobrenome": "Brasil",
	}
	
	for chave, valor := range usuario {
        fmt.Println(chave, valor)
    }
}

// Nao podemos usar for range em structs 
