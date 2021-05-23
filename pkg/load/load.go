package load

import (
	"fmt"
	"runtime"
	"time"

	"github.com/brownhash/golog"
	"github.com/brownhash/loadump/pkg/config"
)

func RunLoad(configuration config.Config) {
	golog.Info("Running load dump")

	systemConfig := configuration.Config

	// ------------------------------------------------------------------------
	// HTTP Load Dump
	// ------------------------------------------------------------------------
	httpConfig := config.MergeConfig(systemConfig, configuration.HTTP.ConfigOverride)

	golog.Info(systemConfig)

	for serviceName, serviceConfig := range configuration.HTTP.Services {
		golog.Info(fmt.Sprintf("Running HTTP Load for: %s", serviceName))

		var TotalReq        int = 0
		var TotalTime       float64 = 0
		var Total5xx        int = 0
		var Total4xx        int = 0
		var TotalMiscErr    int = 0

		TotalStatusMap := make(map[string]int, 0)


		// run untill execution time do not exceeds the time limit
		// after each execution wait for given wait period
		// record the number errors and requests made
		for true {
			start := time.Now()
			err5xx, err4xx, miscErr, statusMap := runHttpLoad(httpConfig.Parallelism, serviceConfig.HTTPAddr, serviceConfig.Method, serviceConfig.Header, serviceConfig.Body)
			elapsed := time.Since(start)

			TotalTime += elapsed.Seconds()

			for status, count := range statusMap {
				if c, ok := TotalStatusMap[status]; ok {
					TotalStatusMap[status] = c + count
				} else {
					TotalStatusMap[status] = count
				}
			}

			if TotalTime < httpConfig.ExecutionMinutes {
				TotalReq        += httpConfig.Parallelism
				Total5xx        += err5xx
				Total4xx        += err4xx
				TotalMiscErr    += miscErr
	
				TotalTime       += time.Duration(httpConfig.WaitPeriod)
				time.Sleep(time.Duration(httpConfig.WaitPeriod))
			} else {
				break
			}
		}
	}
}
