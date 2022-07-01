package utils

import (
	"os/exec"
)

func CheckPS(psname string) (string, error) {
	cmd := exec.Command("pgrep", psname)

	op, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(op), nil
}
