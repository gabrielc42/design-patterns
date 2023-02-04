package main

import "fmt"

func main() {
	manuPrinter := &manufact{}
	osComputer := &os{}
	osComputer.setPrinter(manuPrinter)
	osComputer.print()
	fmt.Println()

	// bridge allows abstract implementations
	// if two OS, two types of Printers
	// then bridge patterns allow for branching similar
	// funcs to classes and objects of similar capacity

	// my example above is one OS and one computer manufacturer
	// like an inverse bridge abstract :P

	// separate out responsibilities into different abstract classes
	// - when parent abstract class defines basic rules,
	// then concrete classes add rules
	// - abstract class that references objects, abstract methods that will be defined
	// in each of the concrete classes
}
