package main

import "fmt"

//array tem tamanho fixo
func main() {
	var meuArray [3]int // array de 3 inteiros
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	//imprime o tamanho do array
	fmt.Println(len(meuArray))

	//imprime o array, len pega o tamanho 3, com -1 resulta em 2
	fmt.Println(len(meuArray) - 1)

	//imprime o ultimo elemento do array
	fmt.Println(meuArray[len(meuArray)-1])

	//imprime o primeiro elemento do array
	fmt.Println(meuArray[0])
	
	//itera sobre o array imprimindo o indice e o valor
	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}


}
