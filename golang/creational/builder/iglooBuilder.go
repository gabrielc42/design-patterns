package main

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (i *iglooBuilder) setWindowType() {
	i.windowType = "ice"
}

func (i *iglooBuilder) setDoorType() {
	i.doorType = "passageway"
}

func (i *iglooBuilder) setNumFloor() {
	i.floor = 1
}

func (i *iglooBuilder) getHouse() house {
	return house{
		doorType:   i.doorType,
		windowType: i.windowType,
		floor:      i.floor,
	}
}
