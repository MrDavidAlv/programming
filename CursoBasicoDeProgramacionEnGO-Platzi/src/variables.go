package main

import "fmt"

func variables() {

	fmt.Println("Hola mundo Go!!!")

	//Declaración de constantes
	const pi float64 = 3.141516
	const pi2 = 3.141516

	fmt.Println("Pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaración variables enteras
	base := 12 //se crea y asigna un valor
	base = 3   //se debio haber creado previamente
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
}
