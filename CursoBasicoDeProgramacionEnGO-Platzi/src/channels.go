package main

import "fmt"

func imprimir(mensaje string, c chan<- string) {
	c <- mensaje
}

func channelsGo() {
	//  c := make(chan string) // Canal sin buffer
	c := make(chan string, 1) // Canal con buffer de tamaño 1

	fmt.Println("Hello")

	go imprimir("Bye", c)

	fmt.Println(<-c)
}
