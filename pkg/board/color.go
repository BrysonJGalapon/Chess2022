package board

import "fmt"

type Color uint8

const (
	WHITE Color = iota
	BLACK
)

func (c Color) String() string {
	switch c {
	case WHITE:
		return "WHITE"
	case BLACK:
		return "BLACK"
	}

	panic(fmt.Sprintf("Unhandled switch case: %d", c))
}

func (c Color) Opposite() Color {
	switch c {
	case WHITE:
		return BLACK
	case BLACK:
		return WHITE
	}

	panic(fmt.Sprintf("Unhandled switch case: %s", c))
}
