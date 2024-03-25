//Array sequencia numerada
// tem um único tipo de dado
// possui um tamanho fixo

//slice(fatia) é uma parte(fatia) de um array
//tem um unico tipo de dado
// tamanho variavel

//mapa busca um valor de acordo com a palavra associada
//pode ser chamado de tabela hash, arrays associativos ou dicionarios

package main

import "fmt"

func main(){
	var x [5]int //array de 5 posições
	x[4] = 80 //atribuindo valor a posição 4
	fmt.Println(x) //imprimindo o array
}