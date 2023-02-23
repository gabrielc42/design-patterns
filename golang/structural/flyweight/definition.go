package main

type Definition struct {
	word       word
	definition string
	page       int
}

func newPerson(definition, wordType string) *Definition {
	word, _ := getWordFactorySingleInstance().getWordByType(wordType)
	return &Definition{
		definition: definition,
		word:       word,
	}
}

func (p *Definition) newLocation(page int) {
	p.page = page
}
