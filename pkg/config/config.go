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

func MergeConfig(config systemConfig, configOverride systemConfigOverride) systemConfigOverride {
	mergedConfig := systemConfigOverride{}

	if configOverride.Parallelism > 0 {
		mergedConfig.Parallelism = configOverride.Parallelism
	} else {
		mergedConfig.Parallelism = config.Parallelism
	}

	if configOverride.WaitPeriod > 0 {
		mergedConfig.WaitPeriod = configOverride.WaitPeriod
	} else {
		mergedConfig.WaitPeriod = config.WaitPeriod
	}

	if configOverride.ExecutionMinutes > 0 {
		mergedConfig.ExecutionMinutes = configOverride.ExecutionMinutes
	} else {
		mergedConfig.ExecutionMinutes = config.ExecutionMinutes
	}

	return mergedConfig
}
