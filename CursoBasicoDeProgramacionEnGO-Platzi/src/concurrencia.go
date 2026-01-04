package main

import (
	"fmt"
	"sync"
	"time"
)

func imprime(texto string) {
	fmt.Println(texto)
}

func imprimeConPuntero(texto string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(texto)
}

func concurreciaNoEficiente() {
	imprime("Hello")
	go imprime("world")

	time.Sleep(100 * time.Millisecond * 1)
}

func concurreciaEficienteSinCanales() {
	var wg sync.WaitGroup
	fmt.Println("Hello")
	wg.Add(1)
	go imprimeConPuntero("world", &wg)
	wg.Wait()
}

func concurreciaConFuncionesAnonimas() {

	concurreciaEficienteSinCanales()

	go func(texto string) {
		fmt.Println(texto)
	}("Adios")

	time.Sleep(100 * time.Millisecond * 1)

}

func concurrenciaEnGo() {
	// Ejecución no eficiente
	fmt.Println("Gorutine no eficiente con timer")
	concurreciaNoEficiente()
	// Ejecución eficiente sin canales
	fmt.Println("Gorutine eficiente usando punteros, sin usar canales")
	concurreciaEficienteSinCanales()
	// Ejecución con funciones anónimas
	fmt.Println("Gorutine con funciones anónimas")
	concurreciaConFuncionesAnonimas()
}
