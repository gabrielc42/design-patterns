// singleton: one object instantiation to a class
// should return same instance whenver multiple goroutines are trying to access that instance

package main

import "fmt"

func main() {
	for i := 0; i < 25; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
