package main

type kaleTopping struct {
	pizza pizza
}

func (k *kaleTopping) getPrice() int {
	pizzaPrice := k.pizza.getPrice()
	return pizzaPrice + 4
}
