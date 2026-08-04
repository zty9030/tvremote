package shell

import (
	"bytes"
	"os/exec"
)

func Execute(command string) (*Result, error) {

	cmd := exec.Command("sh", "-c", command)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	result := &Result{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}

	if err != nil {

		if exitErr, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
		} else {
			result.ExitCode = -1
		}

		return result, err
	}

	result.ExitCode = 0

	return result, nil
}
