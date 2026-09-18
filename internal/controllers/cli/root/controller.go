package root

import (
	"github.com/spf13/cobra"
)

func RootController(cmds ...*cobra.Command) *cobra.Command {
	rootController := &cobra.Command{
		Use:           "telenotify [command]",
		SilenceErrors: true,
		SilenceUsage:  true,
		Short:         "Telenotify notification CLI",
		Example:       "telenotify notify \"Hi, how are you?\"",
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	for _, cmd := range cmds {
		rootController.AddCommand(cmd)
	}

	return rootController
}
