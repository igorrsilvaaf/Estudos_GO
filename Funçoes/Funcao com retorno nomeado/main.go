package main

import "fmt"

func calculos(numero1, numero2 int) (soma int, subtracao int) {
	soma = numero1 + numero2
	subtracao = numero1 - numero2
	return
}

func main() {
	soma, subtracao := calculos(27, 25)
	fmt.Println("O valor da soma é:", soma, "\nO valor da subtracao é:", subtracao)
}
