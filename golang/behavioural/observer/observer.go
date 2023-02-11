package main

type observer interface {
	update(string)
	getId() string
}
