package main

import "fmt"

func buclesFor() {

	//for condicional
	fmt.Println("For condicional")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//for while
	fmt.Println("For while")
	counter := 0
	for counter < 5 {
		fmt.Println(counter)
		counter++
	}
	//for forever
	fmt.Println("For forever")
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever == 5 {
			break
		}
	}
}
