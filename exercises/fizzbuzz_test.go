package exercises

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFizzBuzz(t* testing.T) {
	require.Equal(t, "1", FizzBuzz(1))
	require.Equal(t, "Fizz", FizzBuzz(3))
	require.Equal(t, "Buzz", FizzBuzz(5))
	require.Equal(t, "7", FizzBuzz(7))
	require.Equal(t, "Buzz", FizzBuzz(10))
	require.Equal(t, "FizzBuzz", FizzBuzz(15))
}
