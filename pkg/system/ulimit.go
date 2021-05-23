package system

import (
	"fmt"
	"syscall"

	"github.com/brownhash/golog"
)

func GetRlimit() (syscall.Rlimit, error) {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)

	if err != nil {
		golog.Debug("Error while fetching system rLimit")
	}

	return rLimit, err
}

func CheckLimit(parallelism int) bool {
	limits, err := GetRlimit()

	if err != nil {
		golog.Error("Error while fetching system rLimit")
	}

	if parallelism > int(limits.Cur) {
		golog.Debug(fmt.Sprintf("System Ulimit less than required. Current: %v, Max: %v", limits.Cur, limits.Max))
		return false
	}

	return true
}
