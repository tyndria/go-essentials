package exercises

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvenEndedNumbersCount(t* testing.T) {
	require.Equal(t, EvenEndedNumbersCount(), 3184963)
}
