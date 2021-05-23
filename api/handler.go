package api

import (
	"fmt"

	"github.com/brownhash/golog"
	"github.com/brownhash/loadump/pkg/config"
	"github.com/brownhash/loadump/pkg/load"
)

func NodeHandler(runner, standAlone bool, masterAddr, configFile string, uid uint64) {
	if runner {
		golog.Info("Inititing runner node")
	} else {
		config, err := config.ReadConfig(configFile)

		if err != nil {
			golog.Error(err)
		}

		golog.Debug(fmt.Sprintf("Configuration: %v", config))

		if standAlone {
			golog.Info("Inititing standalone master node")
			load.RunLoad(config)
		} else {
			golog.Info("Inititing master node")
		}
	}
}
