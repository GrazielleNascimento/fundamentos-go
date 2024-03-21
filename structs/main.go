package main

import "fmt"

//Go nao é orientado a objetos, mas tem structs que sao parecidas com classes
// mas nao sao uma classe
//struct é um tipo de dado que pode ser usado para criar tipos de dados mais complexos

//type é usado para criar um novo tipo de dado
//Cliente é o nome do tipo de dado
//struct é a palavra chave para criar um novo tipo de dado
type Cliente struct {
	nome  string
	idade int
	ativo bool
}

func main() {

	grazielle := Cliente{
		nome:  "Grazielle",
		idade: 21,
		ativo: true,
	}

	fmt.Printf(("Nome: %s\n Idade: %d\n Ativo: %t"),grazielle.nome, grazielle.idade, grazielle.ativo)

	grazielle.ativo = false
	fmt.Println()

	fmt.Printf(("\nNome: %s\n Idade: %d\n Ativo: %t"),grazielle.nome, grazielle.idade, grazielle.ativo)

}
