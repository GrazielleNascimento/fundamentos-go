package main

//slice é um array dinamico
//não tem tamanho fixo
import "fmt"

func main() {

	//cria um slice de inteiros com tamanho 5
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	//imprime o slice e o tamanho do slice, len pega o tamanho 5 e cap pega a capacidade do slice
	//%d é um placeholder para inteiros
	//%v é um placeholder para o valor
	//cap é a capacidade do sliceu
	fmt.Printf("\nlen=%d cap=%d  %v", len(s), cap(s), s)

	//cria um slice de inteiros com tamanho 0, apagando todos os elementos do slice anterior
	fmt.Printf("\nlen=%d cap=%d  %v", len(s[:0]), cap(s[:0]), s[:0])

	// os : são usados para fatiar o slice e pegar os elementos do slice de acordo com a posição do slice que você quer pegar
	fmt.Printf("\nlen=%d cap=%d  %v", len(s[:4]), cap(s[:4]), s[:4])

	//ignorei(fatiei) os 2 primeiros elementos do slice,
	//capacidade diminuida por ignorar os 2 primeiros elementos
	fmt.Printf("\nlen=%d cap=%d  %v", len(s[2:]), cap(s[2:]), s[2:])

	//aumenta a capacidade do slice
	// adicionei so um elemento mas é dobrado a capacidade do slice era 10 ficou 20
	s = append(s, 110)

	fmt.Printf("\nlen=%d cap=%d  %v", len(s[:2]), cap(s[:2]), s[:2])
}
