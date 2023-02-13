package main

type cookbook struct {
	book
}

func newCookbook() iBook {
	return &cookbook{
		book: book{
			name:   "Czechoslovak Cookbook",
			author: "Joza Brizova",
		},
	}
}
