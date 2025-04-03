package exercises

import (
	"fmt"
)

type Square struct {
	X      int
	Y      int
	Length int
}

func NewSquare(x, y, length int) (*Square, error) {
	if length < 0 {
		return nil, fmt.Errorf("length should be non-negative: input length is %d", length)
	}

	square := Square{x, y, length}

	return &square, nil
}

func (s *Square) Move(dx, dy int) {
	s.X += dx
	s.Y += dy
}

func (s *Square) Area() int {
	return s.Length * s.Length
}
