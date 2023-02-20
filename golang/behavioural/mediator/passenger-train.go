package main

import "fmt"

type passengerTrain struct {
	mediator mediator
}

func (g *passengerTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("Passenger Train: Landing! 🚉")
	} else {
		fmt.Println("Passenger Train: Waiting... 🎫")
	}
}

func (g *passengerTrain) departure() {
	fmt.Println("Passenger Train: Leaving! 🚝")
	g.mediator.notifyFree()
}

func (g *passengerTrain) permitArrival() {
	fmt.Println("Passenger Train: Arrival Permitted. Landing! 🚧")
}
