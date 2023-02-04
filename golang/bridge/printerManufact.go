package main

import "fmt"

type manufact struct {
}

func (p *manufact) printFile() {
	fmt.Println("Printing by a Manufact Inc. printer")
}
