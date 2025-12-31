package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff [source] [target]",
	Short: "Compare and show differences",
	Long: `Compare two sources and display their differences.
Supports comparing files, directories, configurations, or states.

Example:
  regrada diff file1.txt file2.txt
  regrada diff --context 3
  regrada diff state1 state2 --unified`,
	Args: cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		source := ""
		target := ""

		if len(args) > 0 {
			source = args[0]
		}
		if len(args) > 1 {
			target = args[1]
		}

		context, _ := cmd.Flags().GetInt("context")
		unified, _ := cmd.Flags().GetBool("unified")
		ignoreWhitespace, _ := cmd.Flags().GetBool("ignore-whitespace")

		runDiff(source, target, context, unified, ignoreWhitespace)
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)

	// Add flags specific to diff command
	diffCmd.Flags().IntP("context", "c", 3, "Number of context lines")
	diffCmd.Flags().BoolP("unified", "u", false, "Use unified diff format")
	diffCmd.Flags().BoolP("ignore-whitespace", "w", false, "Ignore whitespace changes")
}

func runDiff(source, target string, context int, unified, ignoreWhitespace bool) {
	fmt.Println("Computing differences...")

	if source != "" && target != "" {
		fmt.Printf("Source: %s\n", source)
		fmt.Printf("Target: %s\n", target)
	} else {
		fmt.Println("Comparing current state with baseline")
	}

	fmt.Printf("Context lines: %d\n", context)

	if unified {
		fmt.Println("Using unified format")
	}

	if ignoreWhitespace {
		fmt.Println("Ignoring whitespace changes")
	}

	// TODO: Implement actual diff logic
	fmt.Println("✓ Diff completed successfully")
}
