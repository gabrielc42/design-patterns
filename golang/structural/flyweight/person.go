package main

type person struct {
	word       word
	personType string
	lat        int
	long       int
}

func newPerson(personType, wordType string) *person {
	word, _ := getWordFactorySingleInstance().getWordByType(wordType)
	return &person{
		personType: personType,
		word:       word,
	}
}

func (p *person) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
