package main

import( 
	"fmt"
)

func main() {
	fmt.Printf("Type: %T - Value: %v\n", true, true)
	fmt.Printf("Tipo: %T - Valor: %v\n", "letras", "letras")
	fmt.Printf("Tipo: %T - Valor: %v\n", 12, 12)
	fmt.Printf("Tipo: %T - Valor: %v\n", "12", "12")
	fmt.Printf("Tipo: %T - Valor: %v\n", 1.345, 1.345)
}

//  %T => Exibe o tipo de variavel
// %v => Exive o valor da variavel