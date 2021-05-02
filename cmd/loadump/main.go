package main

import (
	"fmt"
	"flag"
	"github.com/sharma1612harshit/golog"
	"github.com/sharma1612harshit/loadump/pkg/system"
	"github.com/sharma1612harshit/loadump/pkg/config"
)

func main() {
	logLevel := flag.String("log-level", "INFO", "use -log-level to indicate logging level. DEBUG | INFO | WARNING | ERROR")
	configFile := flag.String("config-file", "config.json", "use -config-file to mention loadump json config file location.")

	flag.Parse()

	golog.SetLogFormat()
	golog.SetLogLevel(*logLevel)

	loadConfig := config.ReadConfig(*configFile)

	system.CheckLimit(loadConfig.Config.Parallelism)

	uid, err := system.GetUid()

	if err != nil {
		golog.Error(fmt.Sprintf("Unable to generate system Uid. %v", err))
	}

	golog.Success(fmt.Sprintf("Successfully generated system Uid: %v", uid))
}