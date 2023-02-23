package main

type redWord struct {
	color string
}

func (c *redWord) getColor() string {
	return c.color
}

func newRedWord() *redWord {
	return &redWord{color: "red"}
}
