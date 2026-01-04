package main

import "fmt"

type figura2D interface {
	area() float64
}

type cuadrado struct {
	lado float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.lado * c.lado
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcularArea(f figura2D) {
	fmt.Println("Area:", f.area())
	// reKturn f.area()
}

func interfacesEnGo() {
	myCuadrado := cuadrado{lado: 4}
	myRectangulo := rectangulo{base: 3, altura: 6}

	// SIN USAR INTERFACES
	fmt.Println("Área del cuadrado:", myCuadrado.area())
	fmt.Println("Área del rectángulo:", myRectangulo.area())

	// USANDO INTERFACES
	calcularArea(myCuadrado)
	calcularArea(myRectangulo)

	// Usando una lista de interfaces
	myInterfaces := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterfaces...)

	// Usando una lista de figuras 2D
	figuras := []figura2D{myCuadrado, myRectangulo}
	for _, figura := range figuras {
		calcularArea(figura)
	}

}
