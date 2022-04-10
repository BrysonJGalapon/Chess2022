package board

import "fmt"

type Color uint8

const (
	WHITE Color = iota
	BLACK
)

func (c *Color) String() string {
	switch *c {
	case WHITE:
		return "WHITE"
	case BLACK:
		return "BLACK"
	}

	panic(fmt.Sprintf("Unhandled switch case: %d", *c))
}
