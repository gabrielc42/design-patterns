// decorator:
// additional things added to an object statically or dynamically,
// enhanced interface to original object

package main

import "fmt"

func main() {
	veggiePizza := &veggie{}

	veggiePizzaExtraCheese := &cheeseTopping{
		pizza: veggiePizza,
	}

	veggiePizzaExtraCheeseExtraKale := &kaleTopping{
		pizza: veggiePizzaExtraCheese,
	}

	fmt.Printf("Price of veggie pizza with extra cheese, extra kale: %d \n", veggiePizzaExtraCheeseExtraKale.getPrice())
}
