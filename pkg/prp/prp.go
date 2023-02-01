package prp

import (
	"os/exec"
)

// logic
const (
	cbd = "brew"
	cbdBundle = "bundle"
	cbdDump = "dump"
)

func CreateBrewDump() (string, error) {
	out, err := exec.Command(cbd, cbdBundle, cbdDump).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}