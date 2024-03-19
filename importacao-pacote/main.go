package main
import "fmt"

const a = "Hello, World"

type ID int
var (
    b bool		= true
    c int		= 10
    d float64 	= 1.2
    e string 	= "Grazi" 
	f ID 		= 1 
)

func main() {
	a := "Short hand so usado na primeira vez q/ a variavel é declarada"
	print("\n", a)
	print("\n", b)
	print("\n", c)
	print("\n", d)
	print("\n", e)
	print("\n", f)
	fmt.Printf("\nO tipo de E, é %T", e) // %T é o tipo da variavel
	fmt.Printf("\nO tipo de E, é %v", e) // %v é o valor da variavel
	fmt.Printf("\nO tipo de F, é %T", f)
}


