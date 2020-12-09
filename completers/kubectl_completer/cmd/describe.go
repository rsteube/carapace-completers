package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Show details of a specific resource or group of resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(describeCmd).Standalone()

	describeCmd.Flags().BoolP("all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is igno")
	describeCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files containing the resource to describe")
	describeCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	describeCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	describeCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	describeCmd.Flags().Bool("show-events", false, "If true, display events related to the described object.")
	rootCmd.AddCommand(describeCmd)

	carapace.Gen(describeCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(""),
		"kustomize": carapace.ActionDirectories(),
	})

	carapace.Gen(describeCmd).PositionalCompletion(
		action.ActionApiResources(),
	)

	carapace.Gen(describeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(args []string) carapace.Action {
			return action.ActionResources("", args[0]).Invoke(args).Filter(args[1:]).ToA()
		}),
	)
}
