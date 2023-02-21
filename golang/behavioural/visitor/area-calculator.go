package main

import "fmt"

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("Calculating area for square! ➕...➖..✖️...➗.. ✔️")
}

func (a *areaCalculator) visitForCircle(c *circle) {
	fmt.Println("Calculating area for circle! ➕...➖..✖️...➗.. ✔️")
}

func (a *areaCalculator) visitForRectangle(r *rectangle) {
	fmt.Println("Calculating area for rectangle! ➕...➖..✖️...➗.. ✔️")
}
