package api

import (
	"fmt"

	"github.com/brownhash/golog"
	"github.com/brownhash/loadump/pkg/config"
	"github.com/brownhash/loadump/pkg/system"
)

func NodeHandler(runner, standAlone bool, masterAddr, configFile string) {

	if runner {
		golog.Warn("Initiating runner node")
		golog.Warn(fmt.Sprintf("Will coordinate with master node at: %v", masterAddr))
	} else {
		golog.Warn("Initiating master node")

		loadConfig := config.ReadConfig(configFile)
		system.CheckLimit(loadConfig.Config.Parallelism)

		if standAlone {
			golog.Warn("Initiating standalone master node")
		}
	}
}
