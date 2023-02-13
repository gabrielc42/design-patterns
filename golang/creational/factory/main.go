package main

import "fmt"

func main() {
	poetry, _ := getBook("poetry")
	cookbook, _ := getBook("cookbook")

	printDetails(poetry)
	printDetails(cookbook)
}

func printDetails(b iBook) {
	fmt.Printf("Book: %s\n", b.getName())
	fmt.Printf("Author: %s\n", b.getAuthor())
}
