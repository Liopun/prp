package api

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

const (
	verifyTokenUrl = "%s/user"
)

type tokenResponseData struct {
	ID        int     `json:"id"`
	Login     string  `json:"login"`
	Name        string     `json:"name"`
	Email        string     `json:"email"`
}

func VerifyToken() (*tokenResponseData, error) {
	response := makeRequest("GET", verifyTokenUrl, nil)
	if response.error != nil {
		return nil, response.error
	}

	jsonData := &tokenResponseData{}
	err := json.Unmarshal(response.body, jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func SetToken(token string) error {
	viper.Set("token", token)

	err := viper.WriteConfig() 
	if err != nil {
		return fmt.Errorf("could not write config: %s", err.Error())
	}

	return nil
}

func SetTokenUser(user, email, name string) error {
	viper.Set("user", user)
	viper.Set("email", email)
	viper.Set("name", name)

	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("could not write config: %s", err.Error())
	}

	return nil
}

func RemoveToken() error {
	viper.Set("token", "")
	viper.Set("user", "")
	viper.Set("email", "")
	viper.Set("name", "")

	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("could not write config: %s", err.Error())
	}

	return nil
}

func IsTokenAvailable() bool {
	return len(viper.GetString("token")) > 0
}

func IsTokenUserAvailable() bool {
	return len(viper.GetString("user")) > 0
}