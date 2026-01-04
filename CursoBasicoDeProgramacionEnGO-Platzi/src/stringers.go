package main

import "fmt"

type laptop struct {
	ram   int
	disk  int
	brand string
}

// Personalizamos el método String() para la estructura laptop
func (myLaptop laptop) String() string {
	return fmt.Sprintf("Laptop: %s, RAM: %dGB, Disk: %dGB", myLaptop.brand, myLaptop.ram, myLaptop.disk)
}

func stringersGo() {
	myLaptop := laptop{ram: 16, disk: 512, brand: "Dell"}
	fmt.Println(myLaptop)

}
