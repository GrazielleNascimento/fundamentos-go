// uma interface define um comportamento de um tipo;
//go não possui excessoes

// A interface Animal define um comportamento que qualquer tipo que deseja ser considerado um "Animal" deve implementar.
type Animal interface {
    // O método Speak retorna a string que representa o som que o animal faz.
    Speak() string
}

// A estrutura Dog representa um cão, que é um tipo de Animal.
type Dog struct {}

// O método Speak para um Dog retorna "Woof!".
func (d Dog) Speak() string {
    return "Woof!"
}

// A estrutura Cat representa um gato, que é um tipo de Animal.
type Cat struct {}

// O método Speak para um Cat retorna "Meow!".
func (c Cat) Speak() string {
    return "Meow!"
}