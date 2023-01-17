package main

import "fmt"

func main() {
	//ARITMETICOS
	soma := 2 + 2
	subtracao := 2 - 2
	divisao := 2 / 2
	multiplicacao := 2 * 2
	restoDivisao := 2 % 2

	fmt.Print(soma, subtracao, divisao, multiplicacao, restoDivisao)

	//Atribuicao
	var variavel string = "ola"
	variavel2 := "ola2"
	fmt.Print(variavel, variavel2)

	//Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//Operadores Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //-> Operador e (AND)
	fmt.Println(verdadeiro || falso) //-> Operador ou (OR)
	fmt.Println(!verdadeiro)         //-> negacao (inverte os valores)

	//Operadores Unarios
	numero1 := 1
	numero1++

	numero1 += 20
	println(numero1)
}
