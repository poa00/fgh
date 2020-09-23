package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure fgh with an interactive prompt",
	Run: func(cmd *cobra.Command, args []string) {
		secretConfig := configure.AskSecretQuestions()
		configure.WriteConfiguration(secretConfig)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}