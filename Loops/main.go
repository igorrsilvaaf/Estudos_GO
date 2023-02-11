package main

import (
	"fmt"
	"time"
)

func main() {
	//tempo := 0
	//
	//for tempo < 10 {
	//	time.Sleep(time.Second)
	//	tempo++
	//	fmt.Println("Tempo", tempo)
	//}
	//fmt.Println(tempo)

	//for tmp := 0; tmp < 10; tmp++ {
	//	time.Sleep(time.Second)
	//	fmt.Println("Segundos", tmp)

	teste := teste(0)
	fmt.Println("Testando", teste)
}

func teste(int) int {
	contador := 0

	for contador < 100 { //enquanto contador for menor que 10
		time.Sleep(time.Second)
		contador += 5
		fmt.Println("Teste", contador)
	}
	return contador
}
