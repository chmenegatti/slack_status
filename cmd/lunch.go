package cmd

import (
	"fmt"
	"os"
	"slack-status/internal/slack"

	"github.com/spf13/cobra"
)

var lunchCmd = &cobra.Command{
	Use:   "lunch",
	Short: "Define status de almoço no Slack",
	Long:  "Define status de almoço no Slack.",
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv(SlackToken)
		api := slack.NewClient(token)

		err := slack.SetLunchTime(api)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao definir status: %v\n", err)
			os.Exit(1)
		}
	},
}
