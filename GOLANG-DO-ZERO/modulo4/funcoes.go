package main

import "fmt"

func main()  {
	fmt.Println(soma(42, 13))
		//ou
	somaValo := soma(42,13)	
	fmt.Println(somaValo)

	sub := subtracao (5, 2)	
	fmt.Println(sub)
//func01
	nome := EscreveNome("MAX")
	fmt.Println(nome)
	
//fun02
 	nome1, nome11, nome21 := EscreveNome1("MAXWELL")
 	fmt.Println(nome1)
 	fmt.Println(nome11)
 	fmt.Println(nome21)

 //fun03
 	nome2, nome12, nome22 := EscreveNome2("MAX","ANNA","ARTHUR")
 	fmt.Println(nome2)
 	fmt.Println(nome12)
 	fmt.Println(nome22)


}
//func01  Obs.: Funções com letra minúscula, elas são privadas e só podem ser utilizadas dentro do pacote "main", com ínicio em letras MAIÚSCULA é uma função pulblica e pode ser utilizada em outros pacotes.
 func EscreveNome (nome string) string {
 	return  nome
}
//fun02
 func EscreveNome1 (nome1 string) (string, string, string) {
  	return  nome1, nome1, nome1
}
 //fun03
  func EscreveNome2 (nome2, nome12, nome22 string) (string, string, string) {
  	return  nome2, nome12, nome22
 }



 func subtracao (x int, y int) int{
 	return x - y 
}
 func soma (x int, y int) int{
 	return x + y 
}