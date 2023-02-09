package main

import "fmt"

func main() {
	numero := 17

	if numero >= 18 {
		fmt.Println("Parabéns voce é de maior")
	} else {
		fmt.Println("Opa voce não pode entrar aqui")
	}

	if novoNumero := numero; novoNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else {
		fmt.Println("Numero é ,emor que zero")
	}

}
