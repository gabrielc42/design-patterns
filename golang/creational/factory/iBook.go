package main

type iBook interface {
	setName(name string)
	setAuthor(author string)
	getName() string
	getAuthor() string
}
