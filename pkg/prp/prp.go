package prp

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

// logic
const (
	cbd = "brew"
	cbdBundle = "bundle"
	cbdDump = "dump"
	cbdRestore = "install"
	brewEnv = "BREW_DIR"
	gitEnv = "GIT_DIR"
)

func CreateBrewDump() (string, error) {
	brewDir := viper.GetString(brewEnv)
	if len(brewDir) == 0 {
		return "", fmt.Errorf("%s dir was not set properly", brewEnv)
	}

	brewFile := brewDir+"/Brewfile"

	_, err := os.Stat(brewFile)
	if err == nil {
		fmt.Println("brew dump file found: getting rid of it first...")
		err = os.Remove(brewFile)
		if err != nil {
			return "", err
		}
	}

	cmd := exec.Command(cbd, cbdBundle, cbdDump)
	cmd.Dir = brewDir

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func RestoreBrewPackages() (string, error) {
	gitDir := viper.GetString(gitEnv)
	if len(gitDir) == 0 {
		return "", fmt.Errorf("%s dir was not set properly", gitEnv)
	}

	brewFile := gitDir+"/Brewfile"
	_, err := os.Stat(brewFile)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(cbd, cbdBundle, cbdRestore)
	cmd.Dir = gitDir

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}