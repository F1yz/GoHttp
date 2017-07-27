package fl1zlog

import (
	"log"
	"os"
)

const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

var level = LevelTrace
var Fl1zLog = log.New(os.Stdout, "", log.LstdFlags | log.Lshortfile)

func SetLevel(l int)  {
	level = l
}

func Level() int {
	return level
}

func Trace(args ...interface{}) {
	if level <= LevelTrace {
		Fl1zLog.Printf("Trace: %v\n", args)
	}
}

func Debug(args ...interface{})  {
	if level <= LevelDebug {
		Fl1zLog.Printf("Debug: %v\n", args)
	}
}

func Info(args ...interface{})  {
	if level <= LevelInfo {
		Fl1zLog.Printf("Info: %v\n", args)
	}
}

func Warning(args ...interface{})  {
	if level <= LevelWarning {
		Fl1zLog.Printf("Warning: %v\n", args)
	}
}

func Error(args ...interface{})  {
	if level <= LevelError {
		Fl1zLog.Printf("Error: %v\n", args)
	}
}

func Critical(args ...interface{})  {
	if level <= LevelCritical {
		Fl1zLog.Printf("Critical: %v\n", args)
	}
}