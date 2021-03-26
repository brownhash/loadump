package logs

import (
	"fmt"
	"log"
	"runtime"
)

// Debug - log debugging messages
func Debug(message interface{}) {
	var fileInfo string = ""

	_, file, lineNum, ok := runtime.Caller(1)

	if ok {
		fileInfo = fmt.Sprintf(SuccessColor, fmt.Sprintf("%s::%d", file, lineNum))
	}

	if GetLogLevel() == "DEBUG" || GetLogLevel() == "" {
		formatter := fmt.Sprintf("%s %s", fileInfo, fmt.Sprintf(DebugColor, message))
		log.Println(formatter)
	}
}

// Info - log informative messages
func Info(message interface{}) {
	if GetLogLevel() == "DEBUG" || GetLogLevel() == "INFO" || GetLogLevel() == "" {
		formatter := fmt.Sprintf(InfoColor, message)
		log.Println(formatter)
	}
}

// Warn - log warning messages
func Warn(message interface{}) {
	if GetLogLevel() == "DEBUG" || GetLogLevel() == "INFO" || GetLogLevel() == "WARN" || GetLogLevel() == "" {
		formatter := fmt.Sprintf(WarningColor, message)
		log.Println(formatter)
	}
}

// Success - log success messages
func Success(message interface{}) {
	if GetLogLevel() == "DEBUG" || GetLogLevel() == "INFO" || GetLogLevel() == "WARN" || GetLogLevel() == "" {
		formatter := fmt.Sprintf(SuccessColor, message)
		log.Println(formatter)
	}
}

// Error - log error messages
func Error(message interface{}) {
	if GetLogLevel() == "DEBUG" || GetLogLevel() == "INFO" || GetLogLevel() == "WARN" || GetLogLevel() == "ERROR" || GetLogLevel() == "" {
		formatter := fmt.Sprintf(ErrorColor, message)
		log.Fatal(formatter)
	}
}