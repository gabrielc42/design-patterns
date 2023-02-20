package main

import "fmt"

type goodsTrain struct {
	mediator mediator
}

func (g *goodsTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("Goods Train: Landing! 🚋")
	} else {
		fmt.Println("Goods Train: Waiting... ⛽")
	}
}

func (g *goodsTrain) departure() {
	g.mediator.notifyFree()
	fmt.Println("Goods Train: Leaving! 🚋")
}

func (g *goodsTrain) permitArrival() {
	fmt.Println("Goods Train: Arrival Permitted. Landing! 🚧")
}
