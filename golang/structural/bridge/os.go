package main

import "fmt"

type os struct {
	printer printer
}

func (o *os) print() {
	fmt.Println("Print request for operating system")
	o.printer.printFile()
}

func (o *os) setPrinter(p printer) {
	o.printer = p
}
