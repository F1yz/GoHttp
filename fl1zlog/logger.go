package fl1zlog

import (
	"log"
	"os"
	"fmt"
)

const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

// not used yet.
type LoggerInterface interface {
	Trace(format string, args ...interface{})

	Debug(format string, args ...interface{})

	Info(format string, args ...interface{})

	Warning(format string, args ...interface{})

	Error(format string, args ...interface{})

	Critical(format string, args ...interface{})
}

var level = LevelTrace
var Fl1zLog = log.New(os.Stdout, "", log.LstdFlags | log.Lshortfile)

func SetLevel(l int)  {
	level = l
}

func Level() int {
	return level
}

func Trace(format string, args ...interface{}) {
	if level <= LevelTrace {
		formatStr := fmt.Sprintf("Trace: %s\n", format)
		Fl1zLog.Printf(formatStr, args)
	}
}

func Debug(format string, args ...interface{})  {
	if level <= LevelDebug {
		formatStr := fmt.Sprintf("Debug: %s\n", format)
		Fl1zLog.Printf(formatStr, args)
	}
}

func Info(format string, args ...interface{})  {
	if level <= LevelInfo {
		formatStr := fmt.Sprintf("Info: %s\n", format)
		Fl1zLog.Printf(formatStr, args)
	}
}

func Warning(format string, args ...interface{})  {
	if level <= LevelWarning {
		formatStr := fmt.Sprintf("Warning: %s\n", format)
		Fl1zLog.Printf(formatStr, args)
	}
}

func Error(format string, args ...interface{})  {
	if level <= LevelError {
		formatStr := fmt.Sprintf("Error: %s\n", format)
		Fl1zLog.Printf(formatStr, args)
	}
}

func Critical(format string, args ...interface{})  {
	if level <= LevelCritical {
		formatStr := fmt.Sprintf("Critical: %s\n", format)
		Fl1zLog.Printf(formatStr, args)
	}
}