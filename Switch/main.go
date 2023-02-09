package main

import "fmt"

// 1 forma de usar Switch
func DiadaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}
}

// 2 forma de usar Switch
func DiadaSemana02(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-feira"
	case numero == 3:
		return "Terça-feira"
	case numero == 4:
		return "Quarta-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sabado"
	default:
		return "Dia da semana inexistente"
	}
}

func DiadaSemana03(numero int) string { //recebe um numero e retorna uma string
	var dia string
	switch {
	case numero == 1:
		dia = "Domingo"
	case numero == 2:
		dia = "Segunda-feira"
	case numero == 3:
		dia = "Terça-feira"
	case numero == 4:
		dia = "Quarta-feira"
	case numero == 5:
		dia = "Quinta-feira"
	case numero == 6:
		dia = "Sexta-feira"
	case numero == 7:
		return "Sabado"
	default:
		return "Esse dia é inexistente"
	}
	return dia
}

func main() {
	//dia := DiadaSemana(5)
	//dia02 := DiadaSemana02(2)
	dia03 := DiadaSemana03(10)
	fmt.Println(dia03)
}
