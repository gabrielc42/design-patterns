package main

type veggie struct{}

func (p *veggie) getPrice() int {
	return 10
}
