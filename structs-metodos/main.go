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
	Endereco // exemplo de composição de struct composta por outra struct
}

//exemplo de metodo para a struct
func (c Cliente) Desativar(){
c.ativo = false
fmt.Printf("O cliente %s foi desativado", c.nome)
}

func main() {

	//criando um novo cliente
	grazielle := Cliente{
		nome:  "Grazielle",
		idade: 21,
		ativo: true,
	}

	//acessando o atributo da struct
	grazielle.ativo = false
	//executando o metodo
	grazielle.Desativar()

	
}
