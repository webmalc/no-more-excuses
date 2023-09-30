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
	configViewer ConfigViewer
}

// configShow show the configuration.
func (s *CommandRouter) configShow(_ *cobra.Command, _ []string) {
	s.configViewer.ShowConfig()
}

// configShow show the configuration.
func (s *CommandRouter) server(_ *cobra.Command, _ []string) {
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
	configViewer ConfigViewer,
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
