package logger

import (
	"flag"
	"log"
)

var flagV = flag.Bool("v", false, "print debug logs")

// Print exposes the print interface for logging
func Print(v ...interface{}) {
	if *flagV {
		log.Println(v...)
	}
}
