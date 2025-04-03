package exercises

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKillServer(t *testing.T) {
	testCases := map[string]struct {
		pidFile     string
		shouldError bool
		nonExisting bool
		content     string
	}{
		"should return error for file path equal empty string": {
			pidFile:     "",
			shouldError: true,
		},
		"should return error for non-existing file": {
			pidFile:     "server-non-existing.pid",
			shouldError: true,
			nonExisting: true,
		},
		"should return error for the file containing empty string": {
			pidFile:     "server-empty.pid",
			shouldError: true,
			content:     "",
		},
		"should return error for the file containing invalid pid": {
			pidFile:     "server-invalid.pid",
			shouldError: true,
			content:     "...",
		},
		"should not return error for the file containing numeric pid": {
			pidFile:     "server.pid",
			shouldError: false,
			content:     "4534",
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			if test.pidFile != "" && !test.nonExisting {
				err := os.WriteFile(test.pidFile, []byte(test.content), 0644)
				if err != nil {
					t.Fatalf("can't create a test file artifact: %v", err)
				}
			}

			err := KillServer(test.pidFile)

			if test.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
