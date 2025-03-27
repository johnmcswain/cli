package cmd

import (
	"fmt"
	"os"

	"github.com/johnmcswain/cli/cmd/git" // Update with your module path if needed.
	"github.com/spf13/cobra"
)

// rootCmd is the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "mcswain",
	Short: "mcswain is a CLI tool that wraps common git commands and more",
	Long:  "mcswain is a command-line tool designed to streamline your workflow with built-in wrappers for git commands like commit, pull, and push.",
}

func init() {
	// Register git commands directly at the top level.
	git.RegisterGitCommands(rootCmd)
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
