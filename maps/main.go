package main

import "fmt"

func main() {
	//mapas são uma coleção de chave e valor
	//mapas são uma coleção desordenada
	//string é a chave e int é o valor
	salarios := map[string]int{"Grazielle": 1200, "João": 1000, "Maria": 500}

	//mostra o mapa
	fmt.Println("Mostra o mapa: ", salarios)

	//mostra o valor da chave Grazielle
	fmt.Println("Mostra o salario de Grazielle: ", salarios["Grazielle"])

	//deleta a chave Maria do mapa
	delete(salarios, "Maria")

	//adiciona um novo valor ao mapa
	salarios["Lara"] = 1500

	//mostra o valor da chave Lara
	fmt.Println("Mostra o salario de Lara: ", salarios["Lara"])

	//mostra o mapa com a chave Maria deletada e a chave Lara adicionada
	fmt.Println("Mostra o mapa: ", salarios)

	//cria um mapa vazio com a função make prepara o mapa para receber valores
	sal := make(map[string]int)

	//mostra o mapa vazio
	fmt.Println("mostra o mapa vazio ", sal)

	//no map tambem pode ser feito assim
	sal1 := map[string]int{}

	sal1["Alice"] = 1000

	fmt.Println("mostra o mapa sal1 Alice ", sal1)

	fmt.Println("mostra o mapa sal1 ", sal1)

	fmt.Println("\nMostra o mapa inteiro com chave e valor: ")

	//itera sobre o mapa imprimindo a chave e o valor
	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	
	fmt.Println("\nMostra o mapa inteiro com valor apenas: ")

	//ignora a chave e imprime so o valor com o underline _ blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
