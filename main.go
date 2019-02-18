package main

import (
	"github.com/RichardKoster/web"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.StandardLogger()

	logger.SetLevel(logrus.DebugLevel)

	w := web.New(":8080")
	go w.Start()

	logrus.Info("Bot is now running.  Press CTRL-C to exit.")
}
