package main

import "fmt"

type aluno struct {
	nome      string
	sobrenome string
	idade     int
	cpf       int
	telefone  int
}

type matricula struct {
	aluno           // Isso é a herança
	numeroMatricula int
	curso           string
	polo            string
}

func main() {
	cadastro := aluno{"Igor", "da Silva Francisco", 27, 10287794940, 48991781573}
	m1 := matricula{cadastro, 01, "Analise e Desenvolvimento de sistema", "Tubarão"}
	fmt.Print(m1)
}
