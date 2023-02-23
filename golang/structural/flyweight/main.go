// flyweight:
// share parts of object state between multiple objects

package main

import "fmt"

func main() {
	game := newGame()
	game.addApple(redType)
	game.addApple(greenType)

	game.addBananas(greenType)
	game.addBananas(yellowType)

	colorFactoryInstance := getColorFactorySingleInstance()

	for colorType, color := range colorFactoryInstance.colorMap {
		fmt.Printf("Fruit Color Type: %s\n Fruit Color: %s\n", colorType, color.getColor())
	}
}
