package main

import (
	"fmt"
)

func switchEnGO() {
	// Switch básico
	day := 3
	switch day {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sábado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Día inválido")
	}

	// Switch con múltiples condiciones
	letter := "a"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vocal")
	default:
		fmt.Println("Consonante")
	}

	// Switch sin expresión
	num := 15
	switch {
	case num < 10:
		fmt.Println("Número menor que 10")
	case num >= 10 && num < 20:
		fmt.Println("Número entre 10 y 19")
	default:
		fmt.Println("Número 20 o mayor")
	}

	// modulo con switch
	number := 10
	switch number % 2 {
	case 0:
		fmt.Println("El número es par")
	case 1:
		fmt.Println("El número es impar")
	}
}
