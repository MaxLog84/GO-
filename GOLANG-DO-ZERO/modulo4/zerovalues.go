package main
import "fmt"
// Apresentar quais valores as variaveis recebem caso nehum valor seja especificado . 
func main()  {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("Inteiro: %v\n", i)
	fmt.Printf("Float: %v\n", f)
	fmt.Printf("Booleano: %v\n", b)
	fmt.Printf("String: %q\n", s)  // %q => serve para imprimir strings entre aspas duplas

}
