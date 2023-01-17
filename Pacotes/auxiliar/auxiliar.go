package auxiliar

import "fmt"

// Escreve uma mensagem na tela
func Escrever() {
	fmt.Println("Escrevendo do pacote aux")
	escrever2() //Declarando a minha funcao escrever2 do pacote auxiliar2
	//Nao ocorreu erro pois eu declarei a funcao no mesmo pacote
	//note que a funcao inicia com letra minuscula
}
