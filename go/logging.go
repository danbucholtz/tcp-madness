package main

import "log"
import "os"

const DEBUG_LEVEL = "debug"
const DEBUG_LEVEL_NUMBER = 1
const WARN_LEVEL = "warn"
const WARN_LEVEL_NUMBER = 2

var loggingNumber = WARN_LEVEL_NUMBER

func InitLogger() {
	levelString := os.Getenv("LOGGING_LEVEL")
	log.Printf("Logger Environment variable String is set to: %s", levelString)
	if len(levelString) == 0 {
		levelString = WARN_LEVEL
	}

	if levelString == DEBUG_LEVEL {
		loggingNumber = DEBUG_LEVEL_NUMBER
	} else {
		loggingNumber = WARN_LEVEL_NUMBER
		levelString = WARN_LEVEL
	}

	log.Printf("Logger level set to: %s", levelString)
}

func Debug(input string) {
	if loggingNumber <= DEBUG_LEVEL_NUMBER {
		log.Println(input)
	}
}

func Debugf(input string, objects ...interface{}) {
	if loggingNumber <= DEBUG_LEVEL_NUMBER {
		log.Printf(input, objects)
	}
}

func Warn(input string) {
	if loggingNumber <= WARN_LEVEL_NUMBER {
		log.Println(input)
	}
}

func Warnf(input string, objects ...interface{}) {
	if loggingNumber <= WARN_LEVEL_NUMBER {
		log.Printf(input, objects)
	}
}
