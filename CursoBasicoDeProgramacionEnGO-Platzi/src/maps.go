package main

import "fmt"

func mapsEnGo() {
	// Declaración de un mapa
	miMapa := make(map[string]int)

	// Agregar elementos al mapa
	miMapa["edad"] = 30
	miMapa["altura"] = 175

	// Imprimir el mapa
	fmt.Println("Mapa:", miMapa)

	// Acceder a un elemento del mapa
	fmt.Println("Edad:", miMapa["edad"])

	// encontrar un valor en el mapa
	valor, existe := miMapa["peso"]
	if existe {
		fmt.Println("Peso:", valor)
	} else {
		fmt.Println("La clave 'peso' no existe en el mapa.")
	}

	// Eliminar un elemento del mapa
	delete(miMapa, "altura")
	fmt.Println("Mapa después de eliminar altura:", miMapa)
	// Recorrer el mapa
	for clave, valor := range miMapa {
		fmt.Println("Clave:", clave, "Valor:", valor)
	}
}
