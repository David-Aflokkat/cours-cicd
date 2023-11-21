package main

import "io"
import "os"

import "gitlab.com/ggpack/logchain-go"
import "gitlab.com/ggpack/webstream"

var streamWriter, wsHandler = webstream.Init()

func initLogging() logchain.Logger {
	params := logchain.Params{
		"template":  "{{.timestamp}} " + version + " {{.levelLetter}} {{.fileLine}} {{.msg}}",
		"verbosity": 3,
		"stream":    io.MultiWriter(os.Stdout, streamWriter),
	}
	chainer := logchain.NewLogChainer(params)

	return chainer.InitLogging()
}

var Logger = initLogging()
