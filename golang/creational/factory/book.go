package main

type book struct {
	name   string
	author string
}

func (b *book) setName(name string) {
	b.name = name
}

func (b *book) getName() string {
	return b.name
}

func (b *book) setAuthor(author string) {
	b.author = author
}

func (b *book) getAuthor() string {
	return b.author
}
