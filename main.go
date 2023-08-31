//go:build !test
// +build !test

package main

import (
	"context"
	"sync"
	"time"
	"webmalc/no-more-excuses/cmd"
	"webmalc/no-more-excuses/common/logger"
)

// TODO: REMOVE TEMP CODE
type TempConfigViewer struct {
	logger *logger.Logger
}

func (s *TempConfigViewer) Run() {
	s.logger.Info("config")
}

type TempServer struct {
	logger *logger.Logger
}

func (s *TempServer) do(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			s.logger.Error("shutdown server!!!!")

			return
		default:
			s.logger.Info("server....")
			time.Sleep(time.Second)
		}
	}
}

func (s *TempServer) Run(ctx context.Context) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(ctx)
	wg.Add(1)
	go s.do(ctx, &wg)
	time.Sleep(time.Second * 10)
	cancel()
	wg.Wait()
}

func main() {
	log := logger.NewLogger()
	router := cmd.NewCommandRouter(
		log, &TempServer{logger: log}, &TempConfigViewer{logger: log},
	)
	router.Run()
}
