package load

import (
	"sync"
	"strings"
	"fmt"

	"github.com/brownhash/golog"
	"github.com/brownhash/loadump/pkg/http"
)

func runHttpLoad(concurrency int, address, reqMethod string, reqHeaders map[string]string, reqBody interface{}) (int, int, int, map[string]int) {
	var errors5xx   int = 0
	var errors4xx   int = 0
	var miscErrors  int = 0

	var waitGroup sync.WaitGroup
	waitGroup.Add(concurrency)

	statusMap := make(map[string]int, 0)

	// run parallel executions equal to the parallelism provided
	for i:=0 ; i<concurrency ; i++ {
		go func(addr, method string, header map[string]string, body interface{}) {
			status, _, err := http.MakeRequest(method, addr, header, body)

			if count, ok := statusMap[status]; ok {
				statusMap[status] = count + 1
			} else {
				statusMap[status] = 1
			}

			if err != nil {
				golog.Info(fmt.Sprintf("%v - %v", status, err))
				miscErrors ++
			}

			if strings.Contains(status, "5") {
				errors5xx ++
			} else if strings.Contains(status, "4") {
				errors4xx ++
			}

			defer waitGroup.Done()

		}(address, reqMethod, reqHeaders, reqBody)
	}

	waitGroup.Wait() // wait until all executions are finished

	return errors5xx, errors4xx, miscErrors, statusMap
}