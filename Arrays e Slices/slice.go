package main

import (
	"fmt"
)

func main() {

	arrayTeste := [...]int{0, 1, 2, 3, 4} //[...] quer dizer que meu array pode ter varias posiçoes
	slice := []int{0, 1, 2, 3, 4}
	fmt.Println(slice, arrayTeste)

	//fmt.Println(reflect.TypeOf(slice)) //TypeOf = Saber o tipo da minha variavel
	//fmt.Println(reflect.TypeOf(arrayTeste))

	slice = append(slice, 5) //adicionando um elemento a mais
	fmt.Println(slice)

	teste := []string{"Igor da Silva"}
	fmt.Println(teste)

	teste = append(teste, "Francisco")
	fmt.Println(teste)

	pedacoArrayTeste := arrayTeste[0:3] //pegando um pedaço do meu array com slice
	fmt.Println(pedacoArrayTeste)

	arrayTeste[1] = 200 //alterando a posiçao 1 do meu array
	fmt.Println(pedacoArrayTeste)
}
