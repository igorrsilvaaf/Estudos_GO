package main

import "fmt"

func main() {

	// usuario := map[string]string{ //dentro do [] é o tipo das chaves e fora dos [] é o tipo dos valores
	// 	"Nome":      " Igor",
	// 	"Sobrenome": " Silva",
	// }
	// fmt.Println(usuario)

	user := map[string]map[string]string{

		"nome": {
			"Nome":      "Igor",
			"Sobrenome": "Silva",
		},
		"Curso": {
			"Nome":   "Analise de Sistemas",
			"Campus": "Tubarão",
		},
	}

	user["logradouro"] = map[string]string{ //Adicionando uma nova chave
		"cidade": "Tubarão",
		"Bairro": "Sertão dos Correias",
		"Estado": "Sc",
	}

	user["contatos"] = map[string]string{
		"Telefone": "48991781573",
		"Email":    "teste@teste.com.br",
	}

	fmt.Println(user)
}
