package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
    case 2:
        return "Segunda"
    case 3:
        return "Terça"
    case 4:
        return "Quarta"
    case 5:
        return "Quinta"
    case 6:
        return "Sexta"
    case 7:	
        return "Sábado"
    default:
        return "Número Inválido"
    }
}

// Segunda maneira, mesmo resultado.
func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch{
		case numero == 1:
			diaDaSemana = "Domingo"
        case numero == 2:
			diaDaSemana = "Segunda"
        case numero == 3:
			diaDaSemana = "Terça"
        case numero == 4:
        	diaDaSemana =  "Quarta"
        case numero == 5:
        	diaDaSemana =  "Quinta"
        case numero == 6:
        	diaDaSemana =  "Sexta"
        case numero == 7:
        	diaDaSemana =  "Sábado"
        default:
            return "Número Inválido"
	}
	return diaDaSemana
}

func main()  {
	fmt.Println("switch")

	dia := diaDaSemana2(3)
	fmt.Println(dia)
}