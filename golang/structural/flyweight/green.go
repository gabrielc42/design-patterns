package main

type green struct {
	color string
}

func (g *green) getColor() string {
	return g.color
}

func newGreen() *green {
	return &green{color: "green"}
}
