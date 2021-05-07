package controller

import "stack/logger"

func TestLogger() {
	logger.PrintLog(TestLogger, "Hey error in controller")
}
