package mcp

import "github.com/spf13/cobra"

func NewMCPRootController(cmds ...*cobra.Command) *cobra.Command {
	rootMCPController := &cobra.Command{
		Use:           "mcp <mcp_type>",
		SilenceErrors: true,
		SilenceUsage:  true,
		Short:         "Telenotify MCP server",
		Example:       "telenotify mcp stdio",
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	for _, cmd := range cmds {
		rootMCPController.AddCommand(cmd)
	}

	return rootMCPController
}
