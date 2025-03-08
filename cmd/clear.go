package cmd

import (
	"fmt"
	"os"
	"slack-status/internal/slack"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Limpa status no Slack",
	Long:  "Limpa status no Slack.",
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv(SlackToken)
		api := slack.NewClient(token)

		err := slack.ClearSlackStatus(api)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao limpar status: %v\n", err)
			os.Exit(1)
		}
	},
}
