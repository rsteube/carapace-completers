package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dmenu",
	Short: "dynamic menu",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("b", false, "dmenu appears at the bottom of the screen.")
	rootCmd.Flags().Bool("f", false, "dmenu grabs the keyboard before reading stdin if not reading from a tty.")
	rootCmd.Flags().String("h", "", "dmenu uses a menu line of at least 'height' pixels tall, but no less than 8.")
	rootCmd.Flags().Bool("i", false, "dmenu matches menu items case insensitively.")
	rootCmd.Flags().String("l", "", "dmenu lists items vertically, with the given number of lines.")
	rootCmd.Flags().String("m", "", "dmenu is displayed on the monitor number supplied.")
	rootCmd.Flags().String("p", "", "defines the prompt to be displayed to the left of the input field.")
	rootCmd.Flags().String("w", "", "sets the width of the dmenu window.")
	rootCmd.Flags().String("x", "", "dmenu is placed at this offset measured from the left side of the monitor.")
	rootCmd.Flags().String("y", "", "dmenu is placed at this offset measured from the top of the monitor.")
	rootCmd.Flags().String("fn", "", "defines the font or font set used.")
	rootCmd.Flags().String("nb", "", "defines the normal background color.")
	rootCmd.Flags().String("nf", "", "defines the normal foreground color.")
	rootCmd.Flags().String("sb", "", "defines the selected background color.")
	rootCmd.Flags().String("sf", "", "defines the selected foreground color.")
	rootCmd.Flags().BoolS("v", "v", false, "prints version information to stdout, then exits.")
}
