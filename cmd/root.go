package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "regrada",
	Short: "Regrada CLI - A powerful command-line tool",
	Long: `Regrada is a CLI tool that provides various commands for managing your workflows.
	
Available commands:
  init  - Initialize a new regrada project
  trace - Trace execution or operations
  diff  - Compare and show differences
  gate  - Manage gates and controls`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags can be added here
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
}

// exitWithError prints an error message and exits
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
