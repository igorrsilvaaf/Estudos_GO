package main

import "fmt"

func inverterSinalComPonteiros(numero *int) {
	*numero = *numero * -1
}

func main() {
	novoNumero := 50
	fmt.Println(novoNumero)
	inverterSinalComPonteiros(&novoNumero)
	fmt.Println(novoNumero)
}
