package main

import "fmt"

func operadoresLogicos() {

	a := true
	b := false
	var c int = 4
	var d int = 5

	fmt.Println("a AND b:", a && b)
	fmt.Println("a OR b:", a || b)
	fmt.Println("NOT a:", !a)

	if a && (2+2 == 4) {
		fmt.Println("a es verdadero y 2+2 es igual a 4")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	if c != d {
		fmt.Println("c es diferente de d")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	if c < d {
		fmt.Println("c es menor que d")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	if c <= d {
		fmt.Println("c es menor o igual que d")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	if c > d {
		fmt.Println("c es mayor que d")
	} else {
		fmt.Println("La condición no se cumplió")
	}

	if c >= d {
		fmt.Println("c es mayor o igual que d")
	} else {
		fmt.Println("La condición no se cumplió")
	}

}
