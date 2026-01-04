package main

import "fmt"

// Ejemplo de uso de defer
func deferExample() {
	defer fmt.Println("Cierrando la conexión a la base de datos...")
	fmt.Println("Usando Defer")
	fmt.Println("Conectando a la base de datos...")
}

// Ejemplo de uso de break y continue
func breakYcontinueEnGo() {
	fmt.Println("Break y continue")
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue
		if i%2 == 0 {
			fmt.Println("Número par:", i)
			continue // Salta a la siguiente iteración si i es par
		}
		// break
		if i == 7 {
			break // Sale del bucle cuando i es 7
		}

		fmt.Println("Número impar:", i)
	}
}

func keywordsEnGo() {
	deferExample()
	breakYcontinueEnGo()
}
