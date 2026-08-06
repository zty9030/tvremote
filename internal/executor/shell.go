package executor

import (
    "fmt"
    "os/exec"
)

type ShellExecutor struct{}

func NewShellExecutor() *ShellExecutor {
    return &ShellExecutor{}
}

func (s *ShellExecutor) Execute(command string) error {

    cmd := exec.Command("sh", "-c", command)

    output, err := cmd.CombinedOutput()

    if err != nil {
        return fmt.Errorf("%v: %s", err, output)
    }

    return nil
}