package exercises

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountWords(t *testing.T) {
	require.Equal(t, 3, CountWords("aaa bb aaa bb bb cc", "bb"))
	require.Equal(t, 2, CountWords("aaa bb aaa BB aa bb cc", "aaa"))
	require.Equal(t, 1, CountWords("aaa bb aaa bb aa bb cc", "cc"))
	require.Equal(t, 1, CountWords("aaa bb aaa bb aa bb CC", "cc"))
	require.Equal(t, 0, CountWords("a", "cc"))
	require.Equal(t, 0, CountWords("a", "c"))
	require.Equal(t, 0, CountWords("", "c"))
	require.Equal(t, 0, CountWords("", ""))
	require.Equal(t, 0, CountWords("aaa bb aaa bb bb cc", ""))
}