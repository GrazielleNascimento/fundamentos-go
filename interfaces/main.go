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
//interface permite que um tipo implemente um metodo
type Pessoa interface{
	Desativar()
}

type Empresa struct {
	Nome string
}

type Cliente struct {
	nome  string
	idade int
	ativo bool
	Endereco // exemplo de composição de struct composta por outra struct
}

func  (e Empresa) Desativar(){
	fmt.Printf("\nA empresa %s foi desativada", e.Nome)
}

//qualquer struct que possui o metodo Desativar implementa a interface Pessoa
//exemplo de metodo para a struct
func (c Cliente) Desativar(){
c.ativo = false
fmt.Printf("\nO cliente %s foi desativado", c.nome)
}

//qualquer struct que possui o metodo Desativar
// implementa a interface Pessoa
func Desativacao(pessoa Pessoa){
	pessoa.Desativar()
}




func main() {

	//criando um novo cliente
	grazielle := Cliente{
		nome:  "Grazielle",
		idade: 21,
		ativo: true,
	}
	minhaEmpresa := Empresa{}
	Desativacao(grazielle) //executando o metodo da interface

	Desativacao(minhaEmpresa) //executando o metodo da interface

	//acessando o atributo da struct
	grazielle.ativo = false
	//executando o metodo
	grazielle.Desativar()

	
}
