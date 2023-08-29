//go:build !test
// +build !test

package main

import (
	"time"
	"webmalc/no-more-excuses/common/logger"
)

func main() {
	l := logger.NewLogger()
	for {
		l.Info("info text")
		l.Error("error text")
		l.Debug("debug text")
		time.Sleep(time.Second)
	}
}
