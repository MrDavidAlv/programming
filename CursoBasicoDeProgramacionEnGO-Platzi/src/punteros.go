package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) upgradeRam(newRam int) {
	myPC.ram = newRam
}

func punterosGo() {
	a := 10
	puntero := &a
	fmt.Println("Valor de a:", a)
	fmt.Println("Dirección de memoria de a:", puntero)
	fmt.Println("Valor al que apunta el puntero:", *puntero)

	*puntero = 20
	fmt.Println("Nuevo valor de a después de modificar a través del puntero:", a)

	myPC := pc{ram: 16, disk: 512, brand: "Dell"}
	fmt.Println("PC antes de modificar:", myPC)

	myPC.ping()
	myPC.upgradeRam(32)
	fmt.Println("PC después de modificar:", myPC)

}
