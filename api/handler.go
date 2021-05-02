package api

import (
	"fmt"
	"github.com/sharma1612harshit/golog"
	"github.com/sharma1612harshit/loadump/pkg/config"
	"github.com/sharma1612harshit/loadump/pkg/system"
)

func NodeHandler(runner, standAlone bool, masterAddr, configFile string) {

	if runner {
		golog.Success("Initiated Loadump runner node")
		golog.Warn(fmt.Sprintf("Will coordinate with master node at: %v", masterAddr))
	} else {
		golog.Success("Initiated Loadump master node")

		loadConfig := config.ReadConfig(configFile)
		system.CheckLimit(loadConfig.Config.Parallelism)

		if standAlone {
			golog.Warn("Standalone master node ...")
		}
	}
}