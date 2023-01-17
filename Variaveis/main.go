package main

import "fmt"

func main() {
	var nome1 string = "Igor da Silva" //declaraçao explicito
	idade1 := 27                       //declaraçao  implicita

	fmt.Println(nome1)
	fmt.Println(idade1)

	var (
		nome  string = "Luana Fernandes"
		idade int    = 25
	)
	fmt.Println(nome, idade)

	teste1, teste2 := "teste1", "teste22"
	fmt.Println(teste1, teste2)

	const constante1 string = "Tipo var, mas nao pode mudar o valor"
	fmt.Println(constante1)

	teste1, teste2 = teste2, teste1 // invertendo os valores das variaveis
	fmt.Println(teste1, teste2)
}
