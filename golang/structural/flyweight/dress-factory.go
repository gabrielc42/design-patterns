package main

import "fmt"

const (
	red   = "red"
	white = "white"
)

var (
	wordFactorySingleInstance = &wordFactory{
		wordMap: make(map[string]word),
	}
)

type wordFactory struct {
	wordMap map[string]word
}

func (w *wordFactory) getWordByType(wordType string) (word, error) {
	if w.wordMap[wordType] != nil {
		return w.wordMap[wordType], nil
	}
	if wordType == red {
		w.wordMap[wordType] = newRedWord()
	}
	if wordType == white {
		w.wordMap[wordType] = newWhiteWord()
		return w.wordMap[wordType], nil
	}
	return nil, fmt.Errorf("Wrong word type passed...")
}

func getWordFactorySingleInstance() *wordFactory {
	return wordFactorySingleInstance
}
