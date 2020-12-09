package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach to a running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachCmd).Standalone()

	attachCmd.Flags().StringP("container", "c", "", "Container name. If omitted, the first container in the pod will be chosen")
	attachCmd.Flags().String("pod-running-timeout", "", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running")
	attachCmd.Flags().BoolP("stdin", "i", false, "Pass stdin to the container")
	attachCmd.Flags().BoolP("tty", "t", false, "Stdin is a TTY")
	rootCmd.AddCommand(attachCmd)

	carapace.Gen(attachCmd).FlagCompletion(carapace.ActionMap{
		"container": carapace.ActionCallback(func(args []string) carapace.Action {
			if len(args) == 0 {
				return carapace.ActionMessage("no resource specified")
			} else {
				return action.ActionContainers("", args[0])
			}
		}),
	})

	carapace.Gen(attachCmd).PositionalCompletion(
		action.ActionApiResourceResources(),
	)
}
