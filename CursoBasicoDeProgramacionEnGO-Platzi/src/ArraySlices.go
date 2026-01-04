package main

import (
	"fmt"
	"strings"
)

func ejemploArraysEnGo() {
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	// longitud del array
	fmt.Println("Longitud del array:", len(array))

	// capacidad del array
	fmt.Println("Capacidad del array:", cap(array))
}

func ejemploSlicesEnGo() {
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Método en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Agregar elementos al slice
	slice = append(slice, 7)
	fmt.Println(slice)

	//append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// Eliminar elementos del slice
	slice = append(slice[:2], slice[4:]...)
	fmt.Println(slice)

	// recorrido del slice
	slice2 := []string{"Go", "es", "un", "lenguaje", "de", "programación"}
	for index := range slice2 {
		fmt.Println("Índice:", index)
	}

	// recorrido del slice con valor
	for index, value := range slice2 {
		fmt.Println("Índice:", index, "Valor:", value)
	}

	// recorrido del slice con valor y índice ignorado
	for _, value := range slice2 {
		fmt.Println("Valor:", value)
	}
}

func esPalindromo(texto string) {
	var textReverse string

	fmt.Println("El texto:", texto)
	// convertir a minúsculas y eliminar espacios
	texto = strings.ToLower(texto)

	for i := len(texto) - 1; i >= 0; i-- {
		textReverse += string(texto[i])
	}

	if textReverse == texto {
		fmt.Println("Es un palíndromo")
	} else {
		fmt.Println("No es un palíndromo")
	}
}

func arrayEnGo() {
	fmt.Println("Arrays son inmutables en tamaño, los slices son dinámicos")
	ejemploArraysEnGo()
	ejemploSlicesEnGo()
	esPalindromo("amor a romA")
}
