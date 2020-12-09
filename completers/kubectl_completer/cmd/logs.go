package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Print the logs for a container in a pod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()

	logsCmd.Flags().Bool("all-containers", false, "Get all containers' logs in the pod(s).")
	logsCmd.Flags().StringP("container", "c", "", "Print the logs of this container")
	logsCmd.Flags().BoolP("follow", "f", false, "Specify if the logs should be streamed.")
	logsCmd.Flags().Bool("ignore-errors", false, "If watching / following pod logs, allow for any errors that occur to be non-fatal")
	logsCmd.Flags().Bool("insecure-skip-tls-verify-backend", false, "Skip verifying the identity of the kubelet that logs are requested from.  In theory, an attacker cou")
	logsCmd.Flags().String("limit-bytes", "", "Maximum bytes of logs to return. Defaults to no limit.")
	logsCmd.Flags().String("max-log-requests", "", "Specify maximum number of concurrent logs to follow when using by a selector. Defaults to 5.")
	logsCmd.Flags().String("pod-running-timeout", "", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running")
	logsCmd.Flags().Bool("prefix", false, "Prefix each log line with the log source (pod name and container name)")
	logsCmd.Flags().BoolP("previous", "p", false, "If true, print the logs for the previous instance of the container in a pod if it exists.")
	logsCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on.")
	logsCmd.Flags().String("since", "", "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one o")
	logsCmd.Flags().String("since-time", "", "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / sin")
	logsCmd.Flags().String("tail", "", "Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwis")
	logsCmd.Flags().Bool("timestamps", false, "Include timestamps on each line in the log output")
	rootCmd.AddCommand(logsCmd)

	carapace.Gen(logsCmd).FlagCompletion(carapace.ActionMap{
		"container": carapace.ActionCallback(func(args []string) carapace.Action {
			if len(args) == 0 {
				return carapace.ActionMessage("no resource specified")
			} else {
				return action.ActionContainers("", args[0])
			}
		}),
	})

	carapace.Gen(logsCmd).PositionalCompletion(
		action.ActionApiResourceResources(),
	)
}
