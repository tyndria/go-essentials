package exercises

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func KillServer(pidFile string) error {
	pidData, err := os.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "can't open pid file")
	}

	fileInfo, err := os.Stat(pidFile)
	if err != nil {
		return errors.Wrap(err, "can't check file size")
	}

	if fileInfo.Size() == 0 {
		return errors.New("file is empty")
	}

	pid, err := strconv.Atoi(string(pidData))
	if err != nil {
		return errors.Wrap(err, "can't convert file data to int")
	}

	fmt.Printf("killing server with pid = %d\n", pid)
	return nil
}
