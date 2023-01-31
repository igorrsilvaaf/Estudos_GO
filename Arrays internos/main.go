package main

import "fmt"

func main() {
	teste := make([]float64, 10, 15)
	fmt.Println(teste)
	fmt.Println(len(teste)) //funcao len mede o tamanho da posicao
	fmt.Println(cap(teste)) //funcao cap para ver a capacidade
}
