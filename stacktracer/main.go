package main

import (
	control "stack/controller"
	logger "stack/logger"
)

/*
Purpose : to print customized echo logs
input : string message to print
input example : logger.Echo("Hey error in main")
output :	FileName : packageName/FunctionName: Line number: : ErrorMessage
Output Example:  main.go : main.main : Line:  11 -> Hey error in main
*/
func main() {
	logger.Echo("Hey error in main")
	control.TestLogger()
}
