package api

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	verifyTokenUrl = "https://api.github.com/issues"
)

func VerifyToken() error {
	response := makeRequest("GET", verifyTokenUrl, nil)
	if response.statusCode != 200 {
		if response.error != nil {
			return response.error
		}

		return fmt.Errorf("%s", response.message)
	}

	return nil
}

func SetToken(token string) error {
	viper.Set("token", token)

	err := viper.WriteConfig() 
	if err != nil {
		return fmt.Errorf("could not write config: %s", err.Error())
	}

	return nil
}

func IsTokenAvailable() bool {
	return len(viper.GetString("token")) > 0
}