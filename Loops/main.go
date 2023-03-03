package main

import (
	"fmt"
	"time"
)

func main() {
	tempo := 0

	for tempo < 10 {time.Sleep(time.Second)
	tempo++
	fmt.Println("Tempo", tempo)
	}
	fmt.Println(tempo)

	for tmp := 0; tmp < 10; tmp++ {
		time.Sleep(time.Second)
		fmt.Println("Segundos", tmp)

	teste := teste(80)
	fmt.Println("Testando", teste)

	teste2 := names
	fmt.Println(teste2)

	//iterar em um ARRAY
	nomes := [3]string{"Igor", "ama", "Luana"}
	for posicao, nome := range nomes {
		fmt.Println(posicao, nome)
	}

	//Iterar em uma MAP
	mapa := map[string]string{
		"Nome":      "Igor",
		"Sobrenome": "Silva",
	}
	for indice, nome := range mapa {
		fmt.Println(indice, nome)
	}
}

//func teste(int) int {
//	contador := 0
//
//	for contador <= 10 { //enquanto contador for menor que 10
//		time.Sleep(time.Second)
//		contador += 1
//		fmt.Println("Teste", contador)
//	}
//	return contador
//}
