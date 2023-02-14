package prp

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// logic
const (
	cbd = "brew"
	cpd = "port"
	cbdBundle = "bundle"
	cpdBundle = "echo"
	cbdDump = "dump"
	cpdDump = "requested"
	cpdConnector = ">"
	cbdRestore = "install"
	cpdRestore = "install"
	cpdRestoreArg = "$(cat Portfile)"
	cbdUpdate = "update"
	cbdUpgrade = "upgrade"
	brewEnv = "BREW_DIR"
	portEnv = "PORT_DIR"
	gitEnv = "GIT_DIR"
	install = "install"
	shMain = "/bin/sh"
	shMainArg = "-c"
	sudo = "sudo"
)

func IsCommandAvailable(name string) bool {
	err := runCommand("", false, "command", "-v", name)

	return err == nil
}

func CreateBrewDump() error {
	brewDir := viper.GetString(brewEnv)
	if len(brewDir) == 0 {
		return fmt.Errorf("%s dir was not set properly", brewEnv)
	}

	brewFile := brewDir+"/Brewfile"

	userInp := confirmPrompt("it's recommended that you update/upgrade your Homebrew packages before backing them up, Do you want to perform update first? [y|n]: ")
	if userInp {
		if err := runCommand("", true, cbd, cbdUpdate); err != nil {
			return err
		}
		if err := runCommand("", true, cbd, cbdUpgrade); err != nil {
			return err
		}
	}

	fmt.Println("generating homebrew dump file...")
	_, err := os.Stat(brewFile)
	if err == nil {
		fmt.Println("brew dump file found: getting rid of it first...")
		err = os.Remove(brewFile)
		if err != nil {
			return err
		}
	}

	return runCommand(brewDir, false, cbd, cbdBundle, cbdDump)
}

func CreatePortDump() error {
	portDir := viper.GetString(portEnv)
	if len(portDir) == 0 {
		return fmt.Errorf("%s dir was not set properly", portEnv)
	}

	portFile := portDir+"/Portfile"
	fmt.Println("generating macports dump file...")

	_, err := os.Stat(portFile)
	if err == nil {
		fmt.Println("macports dump file found: getting rid of it first...")
		err = os.Remove(portFile)
		if err != nil {
			return err
		}
	}

	outFile, err := os.Create(portFile)
	if err != nil {
		return err
	}

	defer outFile.Close()

	return saveCommandToFile(outFile, cpd, cpdBundle, cpdDump)
}

func RestoreBrewPackages() error {
	gitDir := viper.GetString(gitEnv)
	if len(gitDir) == 0 {
		return fmt.Errorf("%s dir was not set properly", gitEnv)
	}

	return runCommand(gitDir, true, cbd, cbdBundle, cbdRestore)
}

func RestorePortsPackages() error {
	gitDir := viper.GetString(gitEnv)
	if len(gitDir) == 0 {
		return fmt.Errorf("%s dir was not set properly", gitEnv)
	}

	_, content, err := getFileContent(gitDir+"/Portfile")
	if err != nil {
		return err
	}
	pkgList := string(content)
	pkgList = strings.TrimSpace(pkgList)
	pkgList = strings.Replace(pkgList, "\n", "", -1)

	return runCommand(gitDir, true, shMain, shMainArg, fmt.Sprintf("%s %s %s %s", sudo, cpd, cpdRestore, pkgList))
}