package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ProcessStop(hasSudo bool, commands []string) (string, error) {
	term := func() *exec.Cmd {
		if hasSudo {
			return exec.Command("sudo", commands...)
		} else {
			return exec.Command(commands[0], commands[1:]...)
		}
	}

	cmd := term()

	op, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}
	return string(op), nil
}
func ProcessStart(hasSudo bool, commands []string) (string, error) {
	term := func() *exec.Cmd {
		if hasSudo {
			return exec.Command("sudo", commands...)
		} else {
			return exec.Command(commands[0], commands[1:]...)
		}
	}

	cmd := term()

	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Start process launched as background process PID : %d", cmd.Process.Pid), nil
}
