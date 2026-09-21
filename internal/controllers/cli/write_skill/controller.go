package writeskill

import (
	"fmt"

	"github.com/spf13/cobra"
)

type skillWriteDeps interface {
	WriteSkill(local bool) error
}

func NewWriteSkillController(deps skillWriteDeps, cmds ...*cobra.Command) *cobra.Command {
	writeSkill := &cobra.Command{
		Use:           "install",
		Short:         "Install a skill",
		SilenceUsage:  false,
		SilenceErrors: false,
		Example:       "telenotify skill install",
		RunE: func(cmd *cobra.Command, args []string) error {
			local, err := cmd.Flags().GetBool("local")
			if err != nil {
				return err
			}

			err = deps.WriteSkill(local)
			if err != nil {
				return err
			}

			fmt.Println("✅ Skill installed")

			return nil
		},
	}

	writeSkill.Flags().Bool(
		"local",
		false,
		"Should install skill locally",
	)

	for _, cmd := range cmds {
		writeSkill.AddCommand(cmd)
	}

	return writeSkill
}
