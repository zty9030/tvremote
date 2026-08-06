package executor

import (
	"fmt"
	"log"
	"os/exec"
)

type ADBExecutor struct {
	Device string
}

func NewADBExecutor(device string) *ADBExecutor {
	return &ADBExecutor{
		Device: device,
	}
}

func (a *ADBExecutor) Execute(command string) error {

	args := []string{
		"-s",
		a.Device,
		"shell",
		command,
	}

	log.Printf("[ADB] adb %v", args)
	cmd := exec.Command("adb", args...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("[ADB] FAILED %s", output)
		return fmt.Errorf("%v: %s", err, string(output))
	}
	log.Printf("[ADB] SUCCESS\n%s", output)
	return nil
}