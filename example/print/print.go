package main

import (
	"github.com/lfhy/log"
)

func main() {
	log.SetLevel(log.LEVEL_ALL)
	log.Debug("Debug日志")
	log.Warn("Warn日志")
	log.Info("Info日志")
	log.Error("Error日志")
}
