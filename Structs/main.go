package main

import (
	"fmt"
)

type user struct { //funciona como se fosse uma classe
	nome            string
	idade           int16
	endereco        logradouro
	dadosAdicionais dados
}

type logradouro struct {
	cidade string
	bairro string
	cep    uint
}

type dados struct {
	email    string
	telefone int
	cpf      int
	rg       int
}

func main() {

	var cadastro user

	cadastro.nome = "Igor da Silva,"
	cadastro.idade = 27
	cadastro.endereco.cidade = "Tubarão,"
	cadastro.endereco.bairro = "Sertão dos Correias,"
	cadastro.endereco.cep = 88780000
	cadastro.dadosAdicionais.telefone = 48991781573
	cadastro.dadosAdicionais.email = "teste.teste@teste"
	cadastro.dadosAdicionais.cpf = 00000000000
	cadastro.dadosAdicionais.rg = 0000000
	fmt.Println("Nome: ", cadastro.nome,
		"\nIdade: ", cadastro.idade,
		"anos,",
		"\nCidade: ", cadastro.endereco.cidade,
		"\nBairro: ", cadastro.endereco.bairro,
		"\nCep: ", cadastro.endereco.cep,
		"\nTelefone", cadastro.dadosAdicionais.telefone,
		"\nEmail: ", cadastro.dadosAdicionais.email,
		"\nCpf: ", cadastro.dadosAdicionais.cpf,
		"\nRg: ", cadastro.dadosAdicionais.rg)

}
