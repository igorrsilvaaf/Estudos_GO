package main

import (
	"fmt"
)

type user struct { //funciona como se fosse uma classe
	nome     string
	idade    int16
	telefone uint
	endereco logradouro
}

type logradouro struct {
	cidade string
	bairro string
	cep    uint
}

func main() {

	var cadastro user

	cadastro.nome = "Igor da Silva,"
	cadastro.idade = 27
	cadastro.telefone = 48991781573
	cadastro.endereco.cidade = "Tubarão,"
	cadastro.endereco.bairro = "Sertão dos Correias,"
	cadastro.endereco.cep = 88780000
	fmt.Println("Nome: ", cadastro.nome, "\nIdade: ", cadastro.idade, "anos,", "\nTelefone", cadastro.telefone, "\nCidade: ", cadastro.endereco.cidade, "\nBairro: ", cadastro.endereco.bairro, "\nCep: ", cadastro.endereco.cep)

}
