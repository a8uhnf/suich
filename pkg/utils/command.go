package utils

import (
	"io"
	"os/exec"
)

// ExecCommand executes an os shell command
func ExecCommand(bfr io.Writer, args ...string) error {
	command := exec.Command(args[0], args[1:]...)

	command.Stdout = bfr
	if err := command.Run(); err != nil {
		return err
	}

	return nil
}
