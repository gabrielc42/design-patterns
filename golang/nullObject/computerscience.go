package main

type computerscience struct {
	numOfProfessors int
}

func (c *computerscience) getNumOfProfessors() int {
	return c.numOfProfessors
}

func (c *computerscience) getName() string {
	return "computerscience"
}
