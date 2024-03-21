package main

import (
	"errors"
	"fmt"
)

//funcao main é a primeira a ser executada
func main() {
	
	//chama a funcao soma e imprime o resultado
	fmt.Println(sum(1, 2))

	//chama a funcao soma e imprime o resultado
	fmt.Println(soma(1, 2))

	//chama a funcao cond e imprime o resultado 
	fmt.Println(cond(51, 3))

	

	//chama a funcao erro e imprime o resultado
	//valor é o resultado da soma
	//err é o erro
	//se err for diferente de nil imprime o erro
	//se err for igual a nil imprime o valor 


	//Na função erro, se a soma dos dois números de entrada for maior ou igual
	// a 50, ela retorna 0 e um erro. No seu caso, 50 + 10 é 60, que é maior que 50,
	// então a função retorna 0 e o erro "Erro! A soma é maior que 50!".
	valor, err := erro(50, 10)
	
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

//Funcoes para serem passadas no main

//funcao soma recebe dois inteiros e retorna um inteiro
func sum(a int, b int) int {
	return a + b
}

//no go as funcoes podem retornar mais de um valor
func soma(a, b int) (int, bool) {
	return a + b, true
}

//funcao cond recebe dois inteiros e retorna um inteiro e um booleano
//se a soma dos dois inteiros for maior ou igual a 50 retorna true
//se a soma dos dois inteiros for menor que 50 retorna false
func cond(a, b int) (int, bool) {
	 if a + b >= 50 {
		return a + b, true
	 }
	 return a + b, false
	 }

//funcao erro recebe dois inteiros e retorna um inteiro e um erro
//se a soma dos dois inteiros for maior ou igual a 50 retorna um erro
//error.New cria um novo erro
//nil é um valor zero para ponteiros, interfaces, mapas, slices, canais e funções	 
func erro(a, b int) (int, error) {
	 if a + b >= 50 {
		return 0, errors.New("Erro! A soma é maior que 50!")
	 }
	 return a + b, nil
	 }

	 //utiliza o ultimo parametro de uma funcao
	 // para retornar um erro ja que no go nao tem try catch