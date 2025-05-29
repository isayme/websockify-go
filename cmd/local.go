package cmd

import (
	"github.com/isayme/websockify-go/cmd/local"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(localCmd)
}

var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Run local of websockify",
	Run: func(cmd *cobra.Command, args []string) {
		local.Run()
	},
}
