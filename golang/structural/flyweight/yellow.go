package main

type yellow struct {
	color string
}

func (y *yellow) getColor() string {
	return y.color
}

func newYellow() *yellow {
	return &yellow{color: "yellow"}
}
