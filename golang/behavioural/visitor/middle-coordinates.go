package main

import "fmt"

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	fmt.Println("Calculating middle point coordinates for square. 📠")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	fmt.Println("Calculating middle point coordinates for circle. 📼")
}

func (a *middleCoordinates) visitForRectangle(r *rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle. 💽")
}
