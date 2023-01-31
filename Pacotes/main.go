package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever() //Nome do pacote + a minha funcao
	//auxiliar.Escrever2()
	erro := checkmail.ValidateFormat("igorprogramacao24@gmail.com")
	fmt.Println(erro)
}
