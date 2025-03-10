package exercises

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestSquareConstructor(t *testing.T) {
	testCases := map[string]struct {
		x int
		y int
		length int
		shouldError bool
	} {
		"should return error if length is negative": {
			x: 1,
			y: 1,
			length: -5,
			shouldError: true,
		},
		"should not return error if length is non-negative": {
			x: 1,
			y: 1,
			length: 5,
		},
		"should not return error if all arguments equal to 0": {
			x: 0,
			y: 0,
			length: 0,
		},
		"should not return error if x and y are negative": {
			x: -1,
			y: -5,
			length: 4,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			square, err := NewSquare(test.x, test.y, test.length)

			if test.shouldError {
				require.Error(t, err)
				require.Nil(t, square)
			} else {
				require.NoError(t, err)
				require.NotNil(t, square)
			}
		})
	}
}

func TestSquareMove(t *testing.T) {
	testCases := map[string]struct {
		dx int
		dy int
		expectedSquareX int
		expectedSquareY int
	} {
		"should move square correctly (positive dx, dy)": {
			dx: 1,
			dy: 1,
			expectedSquareX: 1,
			expectedSquareY: 6,
		},
		"should move square correctly (negative dx, dy)": {
			dx: -50,
			dy: -100,
			expectedSquareX: -50,
			expectedSquareY: -95,
		},
		"should not move square if dx, dy equals zero": {
			dx: 0,
			dy: 0,
			expectedSquareX: 0,
			expectedSquareY: 5,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			square, _ := NewSquare(0, 5, 3)

			square.Move(test.dx, test.dy)

			require.Equal(t, test.expectedSquareX, square.X)
			require.Equal(t, test.expectedSquareY, square.Y)
		})
	}
}

func TestSquareArea(t *testing.T) {
	testCases := map[string]struct {
		square Square
		expectedArea int
	} {
		"should calculate area correctly": {
			square: Square{1, 5, 3},
			expectedArea: 9,
		},
		"should return area equals to zero if length is zero": {
			square: Square{1, 5, 0},
			expectedArea: 0,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expectedArea, test.square.Area())
		})
	}
}