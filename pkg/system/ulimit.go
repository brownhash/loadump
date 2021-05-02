package system

import (
    "github.com/sharma1612harshit/golog"
    "syscall"
	"fmt"
)

func GetRlimit() (syscall.Rlimit, error) {
	var rLimit syscall.Rlimit
    err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)

	if err != nil {
		golog.Debug(err)
	}

	return rLimit, err
}

func CheckLimit(parallelism int) {
	limits, err := GetRlimit()

	if err != nil {
		golog.Error(err)
	}

	if parallelism > int(limits.Cur) {
		golog.Error(fmt.Sprintf("System Ulimit less than required. Current: %v, Max: %v", limits.Cur, limits.Max))
	}
}
