package exercises

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxV(t *testing.T) {
	require.Equal(t, MaxValue([]int{1, 5, 7}), 7)
	require.Equal(t, MaxValue([]int{10, 6, 0}), 10)
	require.Equal(t, MaxValue([]int{10, 6, 500, 0}), 500)
	require.Equal(t, MaxValue([]int{4534}), 4534)
}