package exercises

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestMaxV(t *testing.T) {
	require.Equal(t, 7, MaxValue([]int{1, 5, 7}))
	require.Equal(t, 10, MaxValue([]int{10, 6, 0}))
	require.Equal(t, 500, MaxValue([]int{10, 6, 500, 0}))
	require.Equal(t, 4534, MaxValue([]int{4534}))
}