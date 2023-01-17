package main

import "fmt"

//func somar(n1 int8, n2 int8) int8 {
//	return n1 + n2
//
//}
//
//func main() {
//	soma := somar(10, 20)
//	fmt.Println(soma)

//func main() {
//	var teste = func(idade int64) int64 { // -> declarando uma funcao a uma variavel
//		//fmt.Println(idade)
//		return idade
//	}
//
//	retorno := teste(27)
//	fmt.Println(retorno)
//
//}

func main() {
	resultadoSoma, resultadoSub := calculadora(20, 40)
	fmt.Println(resultadoSoma, resultadoSub)
	
}

// se meu parametro dor do mesmo tio, basta informar uma vez, como no exemplo abaixo
// numero1 e numero2 ambos sao int
// o retorno é preciso informar 2x conforme a quantidade de parametros que voce usar
func calculadora(numero1, numero2 int) (int, int) {
	soma := numero1 + numero2
	sub := numero1 - numero2
	return soma, sub
}
