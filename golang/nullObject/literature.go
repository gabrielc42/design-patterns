package main

type literature struct {
	numOfProfessors int
}

func (c *literature) getNumOfProfessors() int {
	return c.numOfProfessors
}

func (c *literature) getName() string {
	return "mechanical"
}
