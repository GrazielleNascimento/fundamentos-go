package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

// quando se tem ... antes do tipo, significa que a função pode receber mais de um valor
// quando nao sabe a qtd de valores que vou somar
//exemplo de funcao variadica que recebe mais de um valor
// os ... são usados para indicar que a função pode receber mais de um valor
func sum(numeros ...int) int {
	total := 0

	// _ blank identifier para ignorar o indicce
	for _, numero := range numeros {
		total += numero
	}
	return total
}

//funcao variadica em go é uma funcao que pode receber mais de um valor