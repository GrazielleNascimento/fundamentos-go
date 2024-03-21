// struct com composição serve para criar um novo tipo de dado que é composto por outros tipos de dados

package main

import (
	"fmt"
)

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade	 string
	Estado	 string
}

type Cliente struct {
	nome  string
	idade int
	ativo bool
	// Address Endereco  exemplo de criacao de propriedade do tipo Endereco
	Endereco // exemplo de composição de struct composta por outra struct
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

	// grazielle.Address.Cidade = "São Paulo"
	 grazielle.Endereco.Estado = "SP"

	fmt.Println(grazielle.Estado)

	fmt.Printf(("\nNome: %s\n Idade: %d\n Ativo: %t"),grazielle.nome, grazielle.idade, grazielle.ativo)

}
