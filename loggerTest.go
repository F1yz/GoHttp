package main

import (
	logger "fl1zlog"
)

func main()  {
	logger.SetLevel(logger.LevelTrace)
	logger.Trace("haha", "hey, Im fine, fuck u.")
	logger.Debug("WTF")
	logger.Error("only error show!")
	logger.Critical("so and i!")
}