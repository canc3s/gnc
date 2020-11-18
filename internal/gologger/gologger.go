package gologger

import (
	"fmt"
	"log"
)

// Level defines all the available levels we can log at
type Level int

// Available logging levels
const (
	Null Level = iota
	Fatal
	Silent
	Misc
	Info
)

var (
	MaxLevel = Info
)

// log logs the actual message to the screen
func logger(level Level, format string, args ...interface{}) {
	// Don't log if the level is null
	if level == Null {
		return
	}

	if level <= MaxLevel {

		switch level {
		case Fatal:
			log.Fatalf(format, args...)
		case Misc:
			fmt.Printf(format, args...)
		default:
			log.Printf(format, args...)
		}
	}
}

// Infof writes a info message on the screen with the default label
func Infof(format string, args ...interface{}) {
	logger(Info, format, args...)
}

// Printf prints a string on screen without any extra stuff
func Printf(format string, args ...interface{}) {
	logger(Misc, format, args...)
}

// Fatalf exits the program if we encounter a fatal error
func Fatalf(format string, args ...interface{}) {
	logger(Fatal, format, args...)
}