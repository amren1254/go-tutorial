package logger

import (
	"log"
	"runtime"
	"strings"
)

func Echo(errorMessage string) {
	pc, FullFilePath, LineNumber, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	split := strings.Split(FullFilePath, "/")
	lengthOfFilePath := len(split)
	FileName := split[lengthOfFilePath-1]
	log.Println(FileName, ":", funcName, ": Line:", LineNumber, "->", errorMessage)
}
