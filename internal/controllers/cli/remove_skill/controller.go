package removeskill

import (
	"fmt"

	"github.com/spf13/cobra"
)

type skillRemoveDeps interface {
	RemoveSkill(local bool) error
}

func NewRemoveSkillController(deps skillRemoveDeps, cmds ...*cobra.Command) *cobra.Command {
	removeSkill := &cobra.Command{
		Use:           "remove",
		Short:         "Remove a skill",
		SilenceUsage:  false,
		SilenceErrors: false,
		Example:       "telenotify skill remove",
		RunE: func(cmd *cobra.Command, args []string) error {
			local, err := cmd.Flags().GetBool("local")
			if err != nil {
				return err
			}

			err = deps.RemoveSkill(local)
			if err != nil {
				return err
			}

			fmt.Println("✅ Skill removed")

			return nil
		},
	}

	removeSkill.Flags().Bool(
		"local",
		false,
		"Should remove locally installed skill",
	)

	return removeSkill
}
