package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/docker"
	"github.com/spf13/cobra"
)

var image_loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load a image into minikube",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	image_loadCmd.Flags().Bool("daemon", false, "Cache image from docker daemon")
	image_loadCmd.Flags().Bool("overwrite", true, "Overwrite image even if same image:tag name exists")
	image_loadCmd.Flags().Bool("pull", false, "Pull the remote image (no caching)")
	image_loadCmd.Flags().Bool("remote", false, "Cache image from remote registry")
	imageCmd.AddCommand(image_loadCmd)

	carapace.Gen(image_loadCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.CallbackValue, "/") ||
				strings.HasPrefix(c.CallbackValue, ".") {
				return carapace.ActionFiles()
			}
			return docker.ActionRepositoryTags()
		}),
	)
}
