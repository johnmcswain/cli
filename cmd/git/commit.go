package git

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// CommitCmd is the exported command for wrapping git add & commit.
var CommitCmd = &cobra.Command{
	Use:   "commit [path] [message]",
	Short: "Alias for 'git add [path] && git commit -m [message]'",
	Long:  "Adds the specified file or directory and then commits it with the provided commit message.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		message := args[1]

		// Execute 'git add <path>'
		addCmd := exec.Command("git", "add", path)
		addCmd.Stdout = os.Stdout
		addCmd.Stderr = os.Stderr
		if err := addCmd.Run(); err != nil {
			fmt.Printf("Error executing git add: %v\n", err)
			os.Exit(1)
		}

		// Execute 'git commit -m <message>'
		commit := exec.Command("git", "commit", "-m", message)
		commit.Stdout = os.Stdout
		commit.Stderr = os.Stderr
		if err := commit.Run(); err != nil {
			fmt.Printf("Error executing git commit: %v\n", err)
			os.Exit(1)
		}
	},
}
