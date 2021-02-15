package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Verbose simply specifies if you want verbose logging when running commands.
var Verbose bool

// Root command
var rootCmd = &cobra.Command{
	Use:   "makesite",
	Short: "Starter project in go.",
	Long:  `Starter project in go adapted for chess club flyers`,
	Run: func(cmd *cobra.Command, args []string) {
		// Print the usage if no args are passed in :)
		if err := cmd.Usage(); err != nil {
			log.Fatal(err)
		}
	},
}

// Execute a command
func Execute() {
	// Global flags
	// Set author information.
	rootCmd.PersistentFlags().StringP("author", "a", "Chris Barnes (@chrisbarnes2000)", "Author name for copyright attribution")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output mode")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
