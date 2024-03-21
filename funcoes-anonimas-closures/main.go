//funcoes anonimas ou closures são funcoes que não tem nome e podem ser atribuidas a variaveis ou passadas como argumentos para outras funcoes
//podem ser usadas para criar funcoes que retornam outras funcoes

package main

import (
	"fmt"

)



func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	//funcao anonima que utiliza a funcao sum
	total := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	//cria uma variavel total e atribui o valor 0
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
