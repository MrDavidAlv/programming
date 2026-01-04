package main

import "fmt"

// Declaración de una estructura
type Persona struct {
	Nombre string
	Edad   int
	Altura float64
}

func structsGo() {

	// Crear una instancia de la estructura
	p1 := Persona{
		Nombre: "Juan",
		Edad:   28,
		Altura: 1.75,
	}

	// Imprimir la estructura
	fmt.Println("Persona 1:", p1)

	// Modificar campos de la estructura
	var p2 Persona
	p2.Nombre = "María"
	p2.Edad = 30
	p2.Altura = 1.65

	fmt.Println("Persona 2:", p2)
}
