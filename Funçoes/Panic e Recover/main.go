package main

import "fmt"

func recuperandoPrograma() {
	fmt.Println("Solicitando a recuperaçao do sistema!!")
	if r := recover(); r != nil {
		fmt.Println("Sistema estabilizado")
	}
}
func mediaAluno(nota01, nota02 float64) bool { //vai ter um retorno booleano True ou False
	defer recuperandoPrograma()
	media := (nota01 + nota02) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("Erro grave, sistema finalizado.")
}
func main() {
	fmt.Println(mediaAluno(6, 6))
	fmt.Println("Sucesso!")
}
