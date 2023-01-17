package main

import (
	"fmt"
)

type user struct { //funciona como se fosse uma classe
	nome     string
	idade    int16
	telefone int
}

func main() {
	// 1 forma de usar o struct
	var cadastro user
	cadastro.nome = "Igor da Silva" //populando as variaveis onde cadastro é a minha variavel e nome é meu struct(tipo)
	cadastro.idade = 27c
	cadastro.telefone = 48991781573
	fmt.Println(cadastro)

	//2 forma de usar structs
	cadastro2 := user{"Igor", 27, 48991781573}
	fmt.Println(cadastro2)

	//3 forma de usar structs quando se quer usar apenas uma das estruturas
	//por exemplo quero trazer apenas o nome

	cadastro3 := user{nome: "nome"} //infomo o nome : e o parametro, o mesmo poderia ser feito com a idade ou telefone
	fmt.Println(cadastro3)
}
