package main

import "fmt"

const (
	redType    = "apple"
	greenType  = "apple or banana!"
	yellowType = "banana"
)

var (
	colorFactorySingleInstance = &colorFactory{
		colorMap: make(map[string]color),
	}
)

type colorFactory struct {
	colorMap map[string]color
}

func (c *colorFactory) getColorByType(colorType string) (color, error) {
	if c.colorMap[colorType] != nil {
		return c.colorMap[colorType], nil
	}
	if colorType == redType {
		c.colorMap[colorType] = newRed()
	}
	if colorType == greenType {
		c.colorMap[colorType] = newGreen()
		return c.colorMap[colorType], nil
	}
	if colorType == yellowType {
		c.colorMap[colorType] = newYellow()
		return c.colorMap[colorType], nil
	}
	return nil, fmt.Errorf("Wrong word type passed...")
}

func getColorFactorySingleInstance() *colorFactory {
	return colorFactorySingleInstance
}
