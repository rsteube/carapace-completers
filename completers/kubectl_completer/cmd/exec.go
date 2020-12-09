package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command in a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().StringP("container", "c", "", "Container name. If omitted, the first container in the pod will be chosen")
	execCmd.Flags().StringP("filename", "f", "", "to use to exec into the resource")
	execCmd.Flags().String("pod-running-timeout", "", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running")
	execCmd.Flags().BoolP("stdin", "i", false, "Pass stdin to the container")
	execCmd.Flags().BoolP("tty", "t", false, "Stdin is a TTY")
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"container": carapace.ActionCallback(func(args []string) carapace.Action {
			if len(args) == 0 {
				return carapace.ActionMessage("no resource specified")
			} else {
				return action.ActionContainers("", args[0])
			}
		}),
	})

	carapace.Gen(execCmd).PositionalCompletion(
		action.ActionApiResourceResources(),
	)
}
