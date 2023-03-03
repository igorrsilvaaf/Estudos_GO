package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println(texto)
	}("Passando o parametro da funcao anonima")
}
