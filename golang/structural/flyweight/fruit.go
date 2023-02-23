package main

type Fruit struct {
	color  color
	fruit  string
	amount int
}

func newFruit(fruit, colorType string) *Fruit {
	color, _ := getColorFactorySingleInstance().getColorByType(colorType)
	return &Fruit{
		fruit: fruit,
		color: color,
	}
}

func (f *Fruit) newAmount(amount int) {
	f.amount = amount
}
