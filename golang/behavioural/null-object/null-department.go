package main

type nullDepartment struct {
	numOfProfessors int
}

func (c *nullDepartment) getNumOfProfessors() int {
	return 0
}

func (c *nullDepartment) getName() string {
	return "nullDepartment"
}
