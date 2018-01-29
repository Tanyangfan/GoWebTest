package lib

import (
	"log"
	"time"
)

var debug bool

func LogInit(isDebug bool) {
	debug = isDebug
}

func LogD(tag string, message string) {
	if debug {
		doLog("DEBUG", tag, message)
	}
}

func LogFatal(tag string, message string) {
	doLog("Falal", tag, message)
}

func doLog(priority string, tag string, message string) {
	log.Println("================================")
	log.Printf("[%s] tag=%s time=%s\n", priority, tag, time.Now())
	log.Printf("%s\n", message)
	log.Println("================================")
}
