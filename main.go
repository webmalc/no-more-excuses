//go:build !test
// +build !test

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	"webmalc/no-more-excuses/cmd"
	"webmalc/no-more-excuses/common/config"
	"webmalc/no-more-excuses/common/logger"
	"webmalc/no-more-excuses/internal/repositories"
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

type TestConfig struct {
	Apps map[string]map[string]string `mapstructure:"apps"`
}

func main() {
	config.Setup()
	log := logger.NewLogger()
	appRepo := repositories.NewAppRepository(log)
	router := cmd.NewCommandRouter(
		log, &TempServer{logger: log}, &TempConfigViewer{logger: log},
	)
	router.Run()

	// TODO: REMOVE TEMP CODE
	fmt.Println(appRepo.GetApps())
}
