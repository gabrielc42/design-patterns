package main

type game struct {
	apples  []*Fruit
	bananas []*Fruit
}

func newGame() *game {
	return &game{
		apples:  make([]*Fruit, 1),
		bananas: make([]*Fruit, 1),
	}
}

func (g *game) addApple(colorType string) {
	fruit := newFruit("Honeycrisp", colorType)
	g.apples = append(g.apples, fruit)
	return
}

func (g *game) addBananas(colorType string) {
	fruit := newFruit("Plantain", colorType)
	g.bananas = append(g.bananas, fruit)
	return
}
