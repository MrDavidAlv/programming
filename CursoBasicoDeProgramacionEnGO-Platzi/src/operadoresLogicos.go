package main

import (
	"fmt"
	"log"
	"strconv"
)

func esPar(num int) bool {
	return num%2 == 0
}

func compararContrasenas(pass1, pass2 string) bool {
	return pass1 == pass2
}

func operadoresLogicos() {

	a := true
	b := false
	var c int = 4
	var d int = 5

	fmt.Println("a AND b:", a && b)
	fmt.Println("a OR b:", a || b)
	fmt.Println("NOT a:", !a)

	// AND
	if a && (2+2 == 4) {
		fmt.Println("a es verdadero y 2+2 es igual a 4")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	// OR
	if a || b {
		fmt.Println("Al menos una de las condiciones es verdadera")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	// NOT
	if !b {
		fmt.Println("b es falso")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	// DIFERENTE
	if c != d {
		fmt.Println("c es diferente de d")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	// MENOR QUE
	if c < d {
		fmt.Println("c es menor que d")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	// MENOR O IGUAL QUE
	if c <= d {
		fmt.Println("c es menor o igual que d")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	// MAYOR QUE
	if c > d {
		fmt.Println("c es mayor que d")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	// MAYOR O IGUAL QUE
	if c >= d {
		fmt.Println("c es mayor o igual que d")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	// CONVERTIR TEXTOS A NÚMEROS
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("El valor convertido es:", value)

	// USAR LA FUNCIÓN esPar
	if esPar(value) {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	// USAR LA FUNCIÓN compararContrasenas
	pass1 := "miContraseña123"
	pass2 := "miContraseña123"
	if compararContrasenas(pass1, pass2) {
		fmt.Println("Las contraseñas coinciden")
	} else {
		fmt.Println("Las contraseñas no coinciden")
	}
}
