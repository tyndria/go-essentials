package exercises

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKillServer(t *testing.T) {
	testCases := map[string]struct {
		pidFile     string
		shouldError bool
	}{
		"should return error for file path equal empty string": {
			pidFile:     "",
			shouldError: true,
		},
		"should return error for non-existing file": {
			pidFile:     "./assets/server-non-existing.pid",
			shouldError: true,
		},
		"should return error for the file containing empty string": {
			pidFile:     "./assets/server-empty.pid",
			shouldError: true,
		},
		"should return error for the file containing invalid pid": {
			pidFile:     "./assets/server-invalid.pid",
			shouldError: true,
		},
		"should not return error for the file containing numeric pid": {
			pidFile:     "./assets/server.pid",
			shouldError: false,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			err := KillServer(test.pidFile)

			if test.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
