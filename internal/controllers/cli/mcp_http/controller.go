package mcphttp

import (
	"context"

	"github.com/spf13/cobra"
)

type mcpServerRunner interface {
	Run(ctx context.Context) error
}

func MCPHTTPController(serverRunner mcpServerRunner, cmds ...*cobra.Command) *cobra.Command {
	mcpHTTPController := &cobra.Command{
		Use:           "http",
		SilenceErrors: true,
		SilenceUsage:  true,
		Short:         "Start telenotify MCP server over HTTP",
		Example:       "telenotify mcp http",
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return serverRunner.Run(cmd.Context())
		},
	}

	for _, cmd := range cmds {
		cmd.AddCommand(mcpHTTPController)
	}

	return mcpHTTPController
}
