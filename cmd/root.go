package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "mcswain",
	Short: "mcswain is a CLI tool that bundles common git commands and more.",
	Long: `mcswain is a command-line tool designed to streamline your workflow.
It wraps common git commands and can be expanded to include additional functionality.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add the custom commit command.
	rootCmd.AddCommand(commitCmd)
}
