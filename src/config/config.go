package config

import (
	"os"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func ReadConfig(configFilePath string) config {
	jsonFile, err := os.Open(configFilePath)
	defer jsonFile.Close()
	if err != nil {
		errorLog(fmt.Sprintf("Unable to open config file. Error: %v", err))
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		errorLog(fmt.Sprintf("Unable to read config file. Error: %v", err))
	}

	configuration := config{}
	err = json.Unmarshal(byteValue, &configuration)
	if err != nil {
		errorLog(fmt.Sprintf("Unable to parse config file. Error: %v", err))
	}

	return configuration
}