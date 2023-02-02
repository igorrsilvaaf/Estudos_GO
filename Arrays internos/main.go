package main

import (
	"fmt"
)

func main() {
	//teste := make([]float64, 10, 15) // slice do tipo integer que tem um tamanho 10 e capacidade 15
	//fmt.Println(teste)
	//fmt.Println(len(teste)) //funcao len mede o tamanho da posicao
	//fmt.Println(cap(teste)) //funcao cap para ver a capacidade

	nome := make([]float64, 50, 60)
	fmt.Println(nome)

	nome = append(nome, 60)
	fmt.Println(nome)
	fmt.Println(cap(nome))
}
