package main

import "fmt"

func main()  {

	usuario := map[string]string{  // No Map as chaves / valores tem que ser todos o mesmo tipo. Podemos ter chaves inteiro e valores string, mas eles tem que ser coerentes
		"nome": "Pedro",
		"Sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo": "Pedro",
		},
		"curso":{
			"nome": "Engenharia",
			"campus": "Campus 1",
		},

	}
	fmt.Println(usuario2)

	//Apagar chaves do map
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}
	fmt.Println(usuario2)
	
}