package main

import (
	"fmt"
	"flag"
	"github.com/sharma1612harshit/golog"
	"github.com/sharma1612harshit/loadump/pkg/system"
	"github.com/sharma1612harshit/loadump/pkg/config"
)

func main() {
	runner := flag.Bool("runner", false, "use -runner to initiate runner node")

	logLevel := flag.String("log-level", "INFO", "use -log-level to indicate logging level. DEBUG | INFO | WARNING | ERROR")

	standAlone := flag.Bool("stand-alone", false, "use -stand-alone to run loadump on stand-alone machine")
	configFile := flag.String("config-file", "", "use -config-file to mention loadump json config file location.")

	flag.Parse()

	golog.SetLogFormat()
	golog.SetLogLevel(*logLevel)

	uid, err := system.GetUid()

	if err != nil {
		golog.Error(fmt.Sprintf("Unable to generate system Uid. %v", err))
	}

	golog.Success(fmt.Sprintf("Successfully generated system Uid: %v", uid))

	if *runner {
		golog.Success("Initiated Loadump runner node")
	} else {
		golog.Success("Initiated Loadump master node")

		if *standAlone {
			loadConfig := config.ReadConfig(*configFile)
			system.CheckLimit(loadConfig.Config.Parallelism)
		}
	}
}