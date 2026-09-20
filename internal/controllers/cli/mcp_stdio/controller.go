package mcpstdio

import (
	"context"

	"github.com/spf13/cobra"
)

type mcpServerRunner interface {
	Run(ctx context.Context) error
}

func MCPStdIOController(serverRunner mcpServerRunner, cmds ...*cobra.Command) *cobra.Command {
	mcpStdIOController := &cobra.Command{
		Use:           "stdio",
		SilenceErrors: true,
		SilenceUsage:  true,
		Short:         "Start telenotify MCP server over stdio",
		Example:       "telenotify mcp stdio",
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return serverRunner.Run(cmd.Context())
		},
	}

	for _, cmd := range cmds {
		mcpStdIOController.AddCommand(cmd)
	}

	return mcpStdIOController
}
