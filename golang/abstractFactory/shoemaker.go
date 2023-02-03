package main

type shoemaker struct {
}

func (s *shoemaker) makeShoe() iShoe {
	return &shuShoe{
		shoe: shoe{
			logo: "shu",
			size: 100,
		},
	}
}

func (s *shoemaker) makeShort() iShort {
	return &shuShort{
		short: short{
			logo: "shu",
			size: 100,
		},
	}
}
