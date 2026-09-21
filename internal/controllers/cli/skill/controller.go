package skill

import (
	"fmt"

	"github.com/spf13/cobra"
)

type existsChecker interface {
	CheckExists(local bool) (bool, error)
}

func SkillController(existsChecker existsChecker, cmds ...*cobra.Command) *cobra.Command {
	skillRootController := &cobra.Command{
		Use:           "skill [command]",
		SilenceErrors: true,
		SilenceUsage:  true,
		Short:         "Check if Telenotify skill is installed",
		Example:       "telenotify skill",
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			local, err := cmd.Flags().GetBool("local")
			if err != nil {
				return err
			}

			exists, err := existsChecker.CheckExists(local)
			if err != nil {
				return err
			}

			if exists {
				fmt.Println("✅ Skill installed")
			} else {
				fmt.Println("⚠️  Skill is missing")
			}

			return nil
		},
	}

	skillRootController.Flags().Bool(
		"local",
		false,
		"Should install skill locally",
	)

	for _, cmd := range cmds {
		skillRootController.AddCommand(cmd)
	}

	return skillRootController
}
