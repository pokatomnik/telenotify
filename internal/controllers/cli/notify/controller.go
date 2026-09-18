package notify

import (
	"context"
	"time"

	"github.com/pokatomnik/telenotify/internal/entities/notifreq"
	"github.com/spf13/cobra"
)

const defaultTimeout = time.Duration(time.Second * 30)

type Notifier interface {
	Notify(ctx context.Context, msg notifreq.NotificationRequest) error
}

func Notify(notifier Notifier, cmds ...*cobra.Command) *cobra.Command {
	notifyController := &cobra.Command{
		Use:           "notify <message>",
		SilenceErrors: true,
		SilenceUsage:  true,
		Short:         "Send Telegram message",
		Example:       "telenotify notify \"Hi, how are you?\"",
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			duration, err := cmd.Flags().GetDuration("timeout")
			if err != nil {
				return err
			}

			ctx, cancel := context.WithTimeout(cmd.Context(), duration)
			defer cancel()

			text := args[0]

			err = notifier.Notify(ctx, notifreq.NotificationRequest{
				NotificationText: text,
			})
			if err != nil {
				return err
			}

			return nil
		},
	}

	notifyController.Flags().DurationP(
		"timeout",
		"t",
		defaultTimeout,
		"send message request timeout",
	)

	for _, cmd := range cmds {
		notifyController.AddCommand(cmd)
	}

	return notifyController
}
