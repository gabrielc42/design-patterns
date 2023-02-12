package main

type cabinBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newCabinBuilder() *cabinBuilder {
	return &cabinBuilder{}
}

func (c *cabinBuilder) setWindowType() {
	c.windowType = "glass"
}

func (c *cabinBuilder) setDoorType() {
	c.doorType = "wooden door"
}

func (c *cabinBuilder) setNumFloor() {
	c.floor = 2
}

func (c *cabinBuilder) getHouse() house {
	return house{
		doorType:   c.doorType,
		windowType: c.windowType,
		floor:      c.floor,
	}
}
