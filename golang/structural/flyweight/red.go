package main

type red struct {
	color string
}

func (r *red) getColor() string {
	return r.color
}

func newRed() *red {
	return &red{color: "red"}
}
