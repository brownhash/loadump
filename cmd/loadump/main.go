package main

import (
	"flag"
	"github.com/sharma1612harshit/golog"
	"github.com/sharma1612harshit/loadump/pkg/system"
	"github.com/sharma1612harshit/loadump/pkg/config"
)

func main() {
	// c := flag.Bool("c", false, "use -c to start dui controller node")
	logLevel := flag.String("logLevel", "INFO", "use -logLevel to indicate logging level. DEBUG | INFO | WARNING | ERROR")

	flag.Parse()

	golog.SetLogFormat()
	golog.SetLogLevel(*logLevel)

	loadConfig := config.ReadConfig("runnerconfig.json")

	system.CheckLimit(loadConfig.Config.Parallelism)
}