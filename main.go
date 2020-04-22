package main

import (
	"sandy_log/log"
	"time"
)

func init() {
	log.Setup(log.LevelDebug, true, true)
}
func main() {
	for {
		log.Info("1111111")
		log.Debug("2222222")
		time.Sleep(5 * time.Second)
	}
}
