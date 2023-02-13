package main

import "fmt"

func getBook(bookType string) (iBook, error) {
	if bookType == "poetry" {
		return newPoetry(), nil
	}
	if bookType == "cookbook" {
		return newCookbook(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}
