package git

import "github.com/spf13/cobra"

// RegisterGitCommands registers the git-related commands to the provided root command.
func RegisterGitCommands(root *cobra.Command) {
	root.AddCommand(CommitCmd)
	root.AddCommand(PullCmd)
	root.AddCommand(PushCmd)
}
