package exercises

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func KillServer(pidFile string) error {
	file, err := os.Open(pidFile)

	if err != nil {
		return errors.Wrap(err, "can't open pid file")
	}

	defer file.Close()

	defer func() {
		if err = os.Remove(pidFile); err != nil {
			// We could continue running the program even if we can't remove the file
			log.Printf("can't remove pid file %v", err)
		}
	}()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.Wrap(err, "can't parse pid file data")
	}

	fmt.Printf("killing server with pid = %d\n", pid)

	return nil
}
