package main

type poetry struct {
	book
}

func newPoetry() iBook {
	return &poetry{
		book: book{
			name:   "I Live, I Die, I Burn, I Drown",
			author: "Delmira Agustini",
		},
	}
}
