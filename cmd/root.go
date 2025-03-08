package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	SlackToken = "SLACK_TOKEN"
)

var rootCmd = &cobra.Command{
	Use:   "slack-status",
	Short: "CLI para definir status no Slack",
	Long:  "CLI para definir status no Slack.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao executar o comando: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	token := os.Getenv(SlackToken)
	if token == "" {
		fmt.Fprintf(os.Stderr, "É necessário definir a variável de ambiente %s\n", SlackToken)
		os.Exit(1)
	}

	rootCmd.PersistentFlags().String("token", token, "Token de acesso do Slack")
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(clearCmd)
	rootCmd.AddCommand(lunchCmd)
}
