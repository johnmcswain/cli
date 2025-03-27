package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// commitCmd represents the 'commit' command.
var commitCmd = &cobra.Command{
	Use:   "commit [file] [message]",
	Short: "Alias for 'git add [file] && git commit -m [message]'",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		message := args[1]

		// Run 'git add <file>'
		addCmd := exec.Command("git", "add", file)
		addCmd.Stdout = os.Stdout
		addCmd.Stderr = os.Stderr
		if err := addCmd.Run(); err != nil {
			fmt.Println("Error executing git add:", err)
			os.Exit(1)
		}

		// Run 'git commit -m <message>'
		commitCmd := exec.Command("git", "commit", "-m", message)
		commitCmd.Stdout = os.Stdout
		commitCmd.Stderr = os.Stderr
		if err := commitCmd.Run(); err != nil {
			fmt.Println("Error executing git commit:", err)
			os.Exit(1)
		}
	},
}
