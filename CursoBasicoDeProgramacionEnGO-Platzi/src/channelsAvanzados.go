package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func channelsAvanzadosgo() {
	C := make(chan string, 2) // Canal con buffer de tamaño 2

	C <- "Mensaje 1"
	C <- "Mensaje 2"

	fmt.Println(len(C), cap(C))

	// Range y close en canales
	close(C) // Cerrar el canal antes de iterar sobre él
	for mensaje := range C {
		fmt.Println(mensaje)
	}

	// Select con canales
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Mensaje1", email1)
	go message("Mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1:", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2:", m2)
		}
	}
}
