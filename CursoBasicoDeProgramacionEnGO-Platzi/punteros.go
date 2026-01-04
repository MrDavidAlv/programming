package main

import "fmt"

func punterosGo() {
	a := 10
	puntero := &a
	fmt.Println("Valor de a:", a)
	fmt.Println("Dirección de memoria de a:", puntero)
	fmt.Println("Valor al que apunta el puntero:", *puntero)
}
