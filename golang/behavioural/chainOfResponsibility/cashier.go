package main

import "fmt"

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment completed.")
	}
	fmt.Println("Cashier is receiving payment.")
}

func (c *cashier) setNext(next department) {
	c.next = next
}
