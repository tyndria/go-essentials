package exercises

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvenEndedNumbersCount(t* testing.T) {
	require.Equal(t, 3184963, EvenEndedNumbersCount())
}
