package main

import "fmt"

func main()  {
	var nome string
	nome = "max"
	fmt.Println(nome)

	nome = "ana"
	fmt.Println(nome)

	var idade int
	idade = 42
	fmt.Println(idade)
		//ou
	var idad = 25
	fmt.Println(idad)

// declarar multiplas variaveis

	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	var d = true
	fmt.Println(d)
		//ou
	f := "apple"   // := é uma inferencia, ele já entende que f é uma variavel string e que recebe o valor 
	fmt.Println(f)

// Declarar uma variável Constrante, onde o valor não consegue se alterar
	const idadeBento = 4
	idadeBento2 := 5 
	fmt.Println(idadeBento2)

}
	

