package main

import "fmt"

func main() {
	a := 10

	//para acessar o valor da variavel que o ponteiro aponta, é necessário usar o * antes do ponteiro
	var ponteiro *int = &a

	//
	b := &a

	*b = 30

	// *ponteiro = 20

	//ponteiro é um tipo de variavel que armazena o endereço de memoria de outra variavel
	//&a é o endereço de memoria da variavel a
	fmt.Println(&a)

	//mostra o endereço de memoria da variavel a
	fmt.Println(ponteiro)

	//
	fmt.Println(*b)

}
