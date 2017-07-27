package main

import (
	logger "fl1zlog"
)

func main()  {
	logger.SetLevel(logger.LevelTrace)
	logger.Trace("%v", "haha", "hey, Im fine, fuck u.")
	logger.Debug("%v", "WTF")
	logger.Error("%v", "only error show!")
	logger.Critical("%v", "so and i!")
}