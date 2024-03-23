package main

import "fmt"

func main() {
	var nome string = "Grazi"
	var idade int = 21
	var versao float32 = 2.1

	fmt.Println("Olá, sr.", nome, "sua idade é", idade, "e a versão do programa é", versao)

	//sem a necessidade de declarar o tipo da variavel
	var nome1  = "Grazi"
	var idade1 = 21
	var versao1 = 2.1

	fmt.Println("Olá, sr.", nome1, "sua idade é", idade1, "e a versão do programa é", versao1)
}