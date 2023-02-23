package main

type dictionary struct {
	red   []*Definition
	white []*Definition
}

func newDictionary() *dictionary {
	return &dictionary{
		red:   make([]*Definition, 1),
		white: make([]*Definition, 1),
	}
}

func (c *dictionary) addDefintion(wordType string) {
	// dictionary := newDictionary("")
}
