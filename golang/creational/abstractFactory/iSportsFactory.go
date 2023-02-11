package main

import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "shu" {
		return &shoemaker{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}
