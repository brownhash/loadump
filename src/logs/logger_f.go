package logs

import (
	"fmt"
	"log"
	"runtime"
)

// Debugf - log debugging messages in same line
func Debugf(message interface{}) {
	var fileInfo string = ""

	_, file, lineNum, ok := runtime.Caller(1)

	if ok {
		fileInfo = fmt.Sprintf(SuccessColor, fmt.Sprintf("%s::%d", file, lineNum))
	}

	if GetLogLevel() == "DEBUG" || GetLogLevel() == "" {
		formatter := fmt.Sprintf("%s %s", fileInfo, fmt.Sprintf(DebugColor, message))
		log.Printf(formatter)
	}
}

// Infof - log informative messages in same line
func Infof(message interface{}) {
	if GetLogLevel() == "DEBUG" || GetLogLevel() == "INFO" || GetLogLevel() == "" {
		formatter := fmt.Sprintf(InfoColor, message)
		log.Printf(formatter)
	}
}

// Warnf - log warning messages in same line
func Warnf(message interface{}) {
	if GetLogLevel() == "DEBUG" || GetLogLevel() == "INFO" || GetLogLevel() == "WARN" || GetLogLevel() == "" {
		formatter := fmt.Sprintf(WarningColor, message)
		log.Printf(formatter)
	}
}

// Successf - log success messages in same line
func Successf(message interface{}) {
	if GetLogLevel() == "DEBUG" || GetLogLevel() == "INFO" || GetLogLevel() == "WARN" || GetLogLevel() == "" {
		formatter := fmt.Sprintf(SuccessColor, message)
		log.Printf(formatter)
	}
}

// Errorf - log error messages in same line
func Errorf(message interface{}) {
	if GetLogLevel() == "DEBUG" || GetLogLevel() == "INFO" || GetLogLevel() == "WARN" || GetLogLevel() == "ERROR" || GetLogLevel() == "" {
		formatter := fmt.Sprintf(ErrorColor, message)
		log.Fatalf(formatter)
	}
}
