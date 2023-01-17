package main

import (
	"errors"
	"fmt"
)

func main() {
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
