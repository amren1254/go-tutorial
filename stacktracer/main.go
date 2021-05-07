package main

import (
	control "stack/controller"
	logger "stack/logger"
)

func main() {
	//initialize logger
	//get Package.functionName,filename,line number for the functionName
	logger.PrintLog(main, "Hey error in main")
	control.TestLogger()
}
