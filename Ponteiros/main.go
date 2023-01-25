package main

import "fmt"

func main() {
	var teste string
	var memoria *string

	teste = "Igor da Silva"
	memoria = &teste

	fmt.Println(teste, *memoria)
}
