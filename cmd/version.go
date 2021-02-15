package cmd

import (
	"github.com/foize/go.sgr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version for makesite",
	Long:  "All software has versions. This is makesite's",
	Run: func(cmd *cobra.Command, args []string) {
		sgr.Printf("makesite's version is currently [fg-yellow bold] alpha [reset]")
	},
}
