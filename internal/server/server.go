package server

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Server struct {
	logger Logger
	config Config
}

func (s *Server) Run(ctx context.Context) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(ctx)

	wg.Add(1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer func() {
		signal.Stop(interrupt)
		cancel()
	}()
	go func() {
		select {
		case <-interrupt:
			cancel()
		case <-ctx.Done():
		}
		<-interrupt
		os.Exit(1)
	}()

	go s.exec(ctx, &wg)
	wg.Wait()
}

func (s *Server) exec(ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(s.config.IntervalDuration)
	defer wg.Done()
	defer ticker.Stop()

	for {
		s.logger.Debugf("server: run the blocker")
		select {
		case <-ctx.Done():
			s.logger.Error("server: cleanup and shutdown")

			return
		case <-ticker.C:
			continue
		}
	}
}

// NewServer returns a new server object.
func NewServer(logger Logger) *Server {
	return &Server{
		logger: logger,
		config: *NewConfig(),
	}
}
