package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/brownhash/golog"
)

func ReadConfig(configFilePath string) (Config, error) {
	configuration := Config{}

	jsonFile, err := os.Open(configFilePath)
	defer jsonFile.Close()

	if err != nil {
		golog.Debug(fmt.Sprintf("Unable to open config file. Error: %v", err))

		return configuration, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		golog.Debug(fmt.Sprintf("Unable to read config file. Error: %v", err))

		return configuration, err
	}

	err = json.Unmarshal(byteValue, &configuration)

	if err != nil {
		golog.Debug(fmt.Sprintf("Unable to parse config file. Error: %v", err))

		return configuration, err
	}

	return configuration, err
}
