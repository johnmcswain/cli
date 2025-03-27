package git

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// PullCmd is the exported command for wrapping git pull.
var PullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Alias for 'git pull'",
	Long:  "Runs 'git pull' with any additional arguments you provide.",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitArgs := append([]string{"pull"}, args...)
		pull := exec.Command("git", gitArgs...)
		pull.Stdout = os.Stdout
		pull.Stderr = os.Stderr
		if err := pull.Run(); err != nil {
			fmt.Printf("Error executing git pull: %v\n", err)
			os.Exit(1)
		}
	},
}
