package main

import "fmt"

func main() {
	shoeFactory, _ := getSportsFactory("shu")
	shuShoe := shoeFactory.makeShoe()
	shuShort := shoeFactory.makeShort()
	printShoeDetails(shuShoe)
	printShortDetails(shuShort)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShortDetails(s iShort) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
