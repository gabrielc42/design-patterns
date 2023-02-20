package main

type collection interface {
	createInterator() iterator
}
