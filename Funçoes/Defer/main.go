package main

import "fmt"

func prova(nota01, nota02, nota03 float64) bool {
	defer fmt.Println("Verificando a media do aluno...")
	media := (nota01 + nota02 + nota03) / 3

	if media >= 7 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(prova(10, 8, 0))
}
