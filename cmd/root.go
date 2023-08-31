package cmd

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger       ErrorLogger
	rootCmd      *cobra.Command
	config       *Config
	serverRunner ContextRunner
	configViewer Runner
}

// configShow show the configuration.
func (s *CommandRouter) configShow(cmd *cobra.Command, args []string) {
	s.configViewer.Run()
}

// configShow show the configuration.
func (s *CommandRouter) server(cmd *cobra.Command, args []string) {
	s.serverRunner.Run(context.Background())
}

// Run the router.
func (s *CommandRouter) Run() {
	s.rootCmd.AddCommand(
		&cobra.Command{
			Use:   "config",
			Short: "Show the configuration",
			Run:   s.configShow,
		},
		&cobra.Command{
			Use:   "server",
			Short: "Run the server",
			Run:   s.server,
		},
	)
	err := s.rootCmd.Execute()
	if err != nil {
		s.logger.Error(errors.Wrap(err, "root command"))
	}
}

// NewCommandRouter creates a new CommandRouter.
func NewCommandRouter(
	log ErrorLogger,
	serverRunner ContextRunner,
	configViewer Runner,
) CommandRouter {
	config := NewConfig()

	return CommandRouter{
		config:       config,
		logger:       log,
		rootCmd:      &cobra.Command{Use: "no-more-excuses"},
		serverRunner: serverRunner,
		configViewer: configViewer,
	}
}
