package main

import "fmt"

func soma(numeros ...int) int { // ... quer dizer que vai ter varios numeros
	total := 0
	for _, valor := range numeros { // _ quer dizer que eu quero ignorar o indice e pegar só o nome
		total += valor
	}
	return total
}

func main() {
	Resultadosoma := soma(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(Resultadosoma)
}
