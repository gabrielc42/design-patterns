package main

type whiteWord struct {
	color string
}

func (c *whiteWord) getColor() string {
	return c.color
}

func newWhiteWord() *whiteWord {
	return &whiteWord{color: "white"}
}
