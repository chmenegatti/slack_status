package cmd

import (
	"fmt"
	"os"
	"slack-status/internal/slack"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Define status no Slack",
	Long:  "Define status no Slack.",
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv(SlackToken)
		api := slack.NewClient(token)

		err := slack.SetSlackStatus(api)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao definir status: %v\n", err)
			os.Exit(1)
		}
	},
}
