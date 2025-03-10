package exercises

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestContentType(t *testing.T) {
	testCases := map[string]struct {
		inputUrl string
		expectedContentType string
		shouldError bool
	} {
		"should return error for url equals to empty string": {
			inputUrl: "",
			expectedContentType: "",
			shouldError: true,
		},
		"should return error for non-url input": {
			inputUrl: "abcd",
			expectedContentType: "",
			shouldError: true,
		},
		"should return text/html for http://example.com/": {
			inputUrl: "http://example.com/",
			expectedContentType: "text/html",
			shouldError: false,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			actualContentType, err := ContentType(test.inputUrl)

			if test.shouldError {
				require.Error(t, err)
				require.Equal(t, actualContentType, "")
			} else {
				require.NoError(t, err)
				require.Equal(t, actualContentType, test.expectedContentType)
			}
		})
	}

}