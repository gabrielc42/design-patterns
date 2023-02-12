package main

import "fmt"

func main() {
	cabinBuilder := getBuilder("cabin")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(cabinBuilder)
	cabinHouse := director.buildHouse()

	fmt.Printf("Cabin house door type: %s\n", cabinHouse.doorType)
	fmt.Printf("Cabin house window type: %s\n", cabinHouse.windowType)
	fmt.Printf("Cabin house num floor: %d\n", cabinHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("Igloo house door type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo house window type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo house num floor: %d\n", iglooHouse.floor)
}
