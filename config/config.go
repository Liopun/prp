package config

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/spf13/viper"
)

const (
	configDir = ".config/"
	brewDir   = "brew/"
	portDir   = "port/"
	nixDir    = "nix/"
	gitDir    = ".git/"
)

func Init() {
	// brew bundle dir
	if _, err := os.Stat(getDirPath(brewDir)); os.IsNotExist(err) {
		err = os.MkdirAll(getDirPath(brewDir), os.ModeDir|0755)
		if err != nil {
			panic(err)
		}
	}

	// port bundle dir
	if _, err := os.Stat(getDirPath(portDir)); os.IsNotExist(err) {
		err = os.MkdirAll(getDirPath(portDir), os.ModeDir|0755)
		if err != nil {
			panic(err)
		}
	}

	// nix bundle dir
	if _, err := os.Stat(getDirPath(nixDir)); os.IsNotExist(err) {
		err = os.MkdirAll(getDirPath(nixDir), os.ModeDir|0755)
		if err != nil {
			panic(err)
		}
	}

	// config file dir
	if _, err := os.Stat(getDirPath(configDir)); os.IsNotExist(err) {
		err = os.Mkdir(getDirPath(configDir), os.ModeDir|0755)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(getConfigPath()); os.IsNotExist(err) {
		err = os.WriteFile(getConfigPath(), []byte("{}"), 0600)
		if err != nil {
			panic(err)
		}
	}

	viper.SetConfigFile(getConfigPath())
	viper.SetDefault("BASE_URL", "https://api.github.com")
	viper.SetDefault("BREW_DIR", getDirPath(brewDir))
	viper.SetDefault("PORT_DIR", getDirPath(portDir))
	viper.SetDefault("PORT_DIR", getDirPath(portDir))
	viper.SetDefault("NIX_DIR", getDirPath(nixDir))
	viper.SetDefault("REPO_NAME", "prp-backup-repo")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func getConfigPath() string {
	return path.Join(getDirPath(configDir), "prp.json")
}

func getDirPath(dirName string) string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("can't get your home directory.")
		os.Exit(1)
	}

	return path.Join(usr.HomeDir, "/.prp/"+dirName)
}
