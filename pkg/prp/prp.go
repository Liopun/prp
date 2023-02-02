package prp

import (
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

// logic
const (
	cbd = "brew"
	cbdBundle = "bundle"
	cbdDump = "dump"
	brewEnv = "BREW_DIR"
)

func CreateBrewDump() (string, error) {
	brewDir := viper.GetString(brewEnv)

	if len(brewDir) == 0 {
		return "", fmt.Errorf("%s dir was not set properly", brewEnv)
	}

	cmd := exec.Command(cbd, cbdBundle, cbdDump)
	cmd.Dir = brewDir

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}