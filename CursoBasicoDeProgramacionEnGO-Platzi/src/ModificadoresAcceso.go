package main

import (
	pk "CursoBasicoDeProgramacionEnGO-Platzi/src/mypackage"
	"fmt"
)

func modificadoresAcceso() {
	// No se puede acceder a carPrivate desde aquí porque es un tipo no exportado
	// var myCar carPrivate // Esto causaría un error de compilación

	// Sin embargo, podemos acceder a CarPublic porque es un tipo exportado
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2021
	fmt.Println("Carro Público:", myCar)

	// Llamar a la función pública
	pk.PrintMessage()

	// No se puede llamar a printPrivateMessage porque es una función no exportada
	// pk.printPrivateMessage() // Esto causaría un error de compilación
}
