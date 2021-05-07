package logger

import (
	"log"
	"reflect"
	"runtime"
	"strings"
)

func PrintLog(i interface{}, errorMessage string) {
	PackageName, FileName := GetFunctionName(i)
	log.Println(FileName, "-->", PackageName, "-->", errorMessage)
}
func GetFunctionName(i interface{}) (string, string) {
	Name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	File, _ := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).FileLine(10)
	split := strings.Split(File, "/")
	length := len(split)
	FileName := split[length-1]
	return Name, FileName
}
