package main

import (
	"flag"
	"fmt"

	"github.com/brownhash/golog"
	"github.com/brownhash/loadump/api"
	"github.com/brownhash/loadump/pkg/system"
)

func main() {
	runner := flag.Bool("runner", false, "use -runner to initiate runner node")
	masterAddr := flag.String("master-addr", "", "use -master-addr to specify master node address")

	logLevel := flag.String("log-level", "INFO", "use -log-level to indicate logging level. DEBUG | INFO | WARNING | ERROR")

	standAlone := flag.Bool("stand-alone", false, "use -stand-alone to run loadump on stand-alone machine")
	configFile := flag.String("config-file", "config.json", "use -config-file to mention loadump json config file location.")

	flag.Parse()

	golog.SetLogFormat()
	golog.SetLogLevel(*logLevel)

	uid, err := system.GetUid()

	if err != nil {
		golog.Error(fmt.Sprintf("Unable to generate system Uid. %v", err))
	}

	golog.Success(fmt.Sprintf("Successfully generated system Uid: %v", uid))

	api.NodeHandler(*runner, *standAlone, *masterAddr, *configFile)
}
