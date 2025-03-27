package git

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// PushCmd is the exported command for wrapping git push.
var PushCmd = &cobra.Command{
	Use:   "push",
	Short: "Alias for 'git push'",
	Long:  "Runs 'git push' with any additional arguments you provide.",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitArgs := append([]string{"push"}, args...)
		push := exec.Command("git", gitArgs...)
		push.Stdout = os.Stdout
		push.Stderr = os.Stderr
		if err := push.Run(); err != nil {
			fmt.Printf("Error executing git push: %v\n", err)
			os.Exit(1)
		}
	},
}
